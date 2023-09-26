package product

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAddLogic) ProductAdd(req *types.AddProductReq) (resp *types.AddProductResp, err error) {
	var AttributeValueList []*pmsclient.AttributeValueList
	for _, i := range req.AttributeValueList {
		AttributeValueList = append(AttributeValueList, &pmsclient.AttributeValueList{
			AttributeID: i.AttributeID,
			Value:       i.Value,
		})
	}
	//
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if ii == "商家" {
			isSJ = true
		}
	}
	merchantID := int64(0)
	if isSJ {
		merchantID = merchant.ID
	}
	//
	_, err = l.svcCtx.Pms.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
		CategoryID:          req.CategoryID,
		AttributeCategoryID: req.AttributeCategoryID,
		Name:                req.Name,
		Pic:                 req.Pic,
		ProductSn:           req.ProductSn,
		Price:               req.Price,
		Desc:                req.Desc,
		OriginalPrice:       req.OriginalPrice,
		Unit:                req.Unit,
		AttributeValueList:  AttributeValueList,
		ImgUrl:              req.ImgUrl,
		IntroduceImgUrl:     req.IntroduceImgUrl,
		MerchantID:          merchantID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddProductResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
