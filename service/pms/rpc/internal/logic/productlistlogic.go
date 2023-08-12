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
		//获取商品的所有属性【】id
		var attrValueIDs []int64
		attrlist, _ := l.svcCtx.SpuAttributeModel.GetAttributeValueByProductID(int64(i.ID))
		for _, i := range attrlist {
			attrValueIDs = append(attrValueIDs, i.AttributeValueID)
		}
		//获取商品下的size
		var sizeList []*pms.SizeList
		size, _ := l.svcCtx.SpuSizeModel.GetSizeBySpuID(int64(i.ID))
		for _, j := range size {
			var sizeValue []*pms.SizeValue
			info, _ := l.svcCtx.SpuSizeValueModel.GetSizeValueBySizeID(int64(j.ID))
			for _, f := range info {
				sizeValue = append(sizeValue, &pms.SizeValue{
					ID:     int64(f.ID),
					SizeID: f.SizeID,
					Value:  f.Value,
				})
			}
			sizeList = append(sizeList, &pms.SizeList{
				ID:        int64(j.ID),
				ProductID: j.ProductID,
				Name:      j.Name,
				SizeValue: sizeValue,
			})
		}
		list = append(list, &pms.ProductListData{
			Id:               int64(i.ID),
			CategoryID:       i.CategoryID,
			Name:             i.Name,
			Pic:              i.Pic,
			ProductSn:        i.ProductSn,
			SubTitle:         i.SubTitle,
			Description:      i.Description,
			OriginalPrice:    i.OriginalPrice,
			Stock:            i.Stock,
			Unit:             i.Unit,
			Sale:             i.Sale,
			Price:            i.Price,
			SizeList:         sizeList,
			AttributeValueID: attrValueIDs,
		})
	}
	return &pms.ProductListResp{
		Total: total,
		List:  list,
	}, nil
}
