package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品列表
func (l *ProductListLogic) ProductList(in *pms.ProductListReq) (*pms.ProductListResp, error) {
	all, total, err := l.svcCtx.ProductModel.GetProductList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.ProductListData
	for _, i := range all {
		list = append(list, &pms.ProductListData{
			Id:                  int64(i.ID),
			CategoryID:          i.CategoryID,
			Name:                i.Name,
			Pic:                 i.Pic,
			ProductSn:           i.ProductSn,
			Desc:                i.Desc,
			OriginalPrice:       i.OriginalPrice,
			Unit:                i.Unit,
			Price:               i.Price,
			AttributeCategoryID: i.AttributeCategoryID,
		})
	}

	return &pms.ProductListResp{
		Total: total,
		List:  list,
	}, nil
}

type AttributeValue struct {
	AttributeID int      `json:"attributeID"`
	Value       []string `json:"value"`
}
