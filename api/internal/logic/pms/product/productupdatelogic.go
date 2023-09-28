package product

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductUpdateLogic) ProductUpdate(req *types.UpdateProductReq) (resp *types.UpdateProductResp, err error) {
	var AttributeValueList []*pmsclient.AttributeValueList
	for _, i := range req.AttributeValueList {
		AttributeValueList = append(AttributeValueList, &pmsclient.AttributeValueList{
			AttributeID: i.AttributeID,
			Value:       i.Value,
		})
	}
	_, err = l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:                  req.Id,
		AttributeCategoryID: req.AttributeCategoryID,
		CategoryID:          req.CategoryID,
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
		MerchantID:          req.MerchantID,
	})
	if err != nil {
		return &types.UpdateProductResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateProductResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
