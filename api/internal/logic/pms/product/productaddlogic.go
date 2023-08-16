package product

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

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
	var size []*pmsclient.Size
	for _, i := range req.Size {
		size = append(size, &pmsclient.Size{
			Name:          i.Name,
			SizeValueName: i.SizeValue,
		})
	}
	_, err = l.svcCtx.Pms.ProductAdd(l.ctx, &pmsclient.ProductAddReq{
		CategoryID:       req.CategoryID,
		Name:             req.Name,
		Pic:              req.Pic,
		ProductSn:        req.ProductSn,
		Price:            req.Price,
		Desc:             req.Desc,
		OriginalPrice:    req.OriginalPrice,
		Unit:             req.Unit,
		AttributeValueID: req.AttributeValueID,
		Size:             size,
		ImgUrl:           req.ImgUrl,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddProductResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
