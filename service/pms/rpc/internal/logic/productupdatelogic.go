package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *ProductUpdateLogic) ProductUpdate(in *pms.ProductUpdateReq) (*pms.ProductUpdateResp, error) {
	err := l.svcCtx.ProductModel.UpdateProduct(in.Id, &model.Product{
		CategoryID:    in.CategoryID,
		Name:          in.Name,
		Pic:           in.Pic,
		ProductSn:     in.ProductSn,
		SubTitle:      in.SubTitle,
		Description:   in.Description,
		OriginalPrice: in.OriginalPrice,
		Stock:         in.Stock,
		Unit:          in.Unit,
		Sale:          in.Sale,
		Price:         in.Price,
	})
	//属性
	_ = l.svcCtx.SpuAttributeModel.DeleteSpuAttributeBySpuID(in.Id)
	for _, i := range in.AttributeValueID {
		_, _ = l.svcCtx.SpuAttributeModel.AddSpuAttribute(&model.SpuAttribute{
			ProductID:        in.Id,
			AttributeValueID: i,
		})
	}
	_ = l.svcCtx.SpuSizeModel.DeleteSpuSizeBySpuID(in.Id)
	for _, i := range in.Size {
		size, _ := l.svcCtx.SpuSizeModel.AddSpuSize(&model.SpuSize{
			ProductID: in.Id,
			Name:      i.Name,
		})
		_ = l.svcCtx.SpuSizeValueModel.DeleteSpuSizeValueBySizeID(i.ID)
		for _, j := range i.SizeValue {
			l.svcCtx.SpuSizeValueModel.AddSpuSizeValue(&model.SpuSizeValue{
				SizeID: int64(size.ID),
				Value:  j.Value,
			},
			)
		}
	}
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &pms.ProductUpdateResp{}, nil
}
