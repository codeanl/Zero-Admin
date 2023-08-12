package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商品
func (l *ProductAddLogic) ProductAdd(in *pms.ProductAddReq) (*pms.ProductAddResp, error) {
	info := model.Product{
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
	}
	spu, err := l.svcCtx.ProductModel.AddProduct(&info)
	//存平台属性 （SpuID,AttributeValueID）spu对应许多个平台属性
	for _, i := range in.AttributeValueID {
		_, _ = l.svcCtx.SpuAttributeModel.AddSpuAttribute(&model.SpuAttribute{
			ProductID:        int64(spu.ID),
			AttributeValueID: i,
		})
	}
	//存规格 SpuID Size->SIzeID SizeValue 两个表
	for _, i := range in.Size {
		size, _ := l.svcCtx.SpuSizeModel.AddSpuSize(&model.SpuSize{
			ProductID: int64(spu.ID),
			Name:      i.Name,
		})
		for _, j := range i.SizeValue {
			l.svcCtx.SpuSizeValueModel.AddSpuSizeValue(&model.SpuSizeValue{
				SizeID: int64(size.ID),
				Value:  j.Value,
			},
			)
		}
	}
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &pms.ProductAddResp{}, nil
}
