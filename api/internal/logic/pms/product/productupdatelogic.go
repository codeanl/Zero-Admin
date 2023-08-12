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
	var size []*pmsclient.Size
	for _, i := range req.Size {
		var sizeValue []*pmsclient.SizeValue
		for _, u := range i.SizeValue {
			sizeValue = append(sizeValue, &pmsclient.SizeValue{
				SizeID: u.SizeID,
				Value:  u.Value,
			})
		}
		size = append(size, &pmsclient.Size{
			ID:        i.ID,
			ProductID: i.ProductID,
			Name:      i.Name,
			SizeValue: sizeValue,
		})
	}
	_, err = l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		Id:               req.Id,
		CategoryID:       req.CategoryID,
		Name:             req.Name,
		Pic:              req.Pic,
		ProductSn:        req.ProductSn,
		Sale:             0,
		Price:            req.Price,
		SubTitle:         req.SubTitle,
		Description:      req.Description,
		OriginalPrice:    req.OriginalPrice,
		Stock:            req.Stock,
		Unit:             req.Unit,
		AttributeValueID: req.AttributeValueID,
		Size:             size,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateProductResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
