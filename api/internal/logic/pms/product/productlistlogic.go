package product

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ListProductReq) (*types.ListProductResp, error) {
	resp, err := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListProductData
	for _, item := range resp.List {
		var sizeList []types.SizeList
		for _, i := range item.SizeList {
			var sizeValueList []types.SizeValueList
			for _, j := range i.SizeValue {
				sizeValueList = append(sizeValueList, types.SizeValueList{
					ID:     j.ID,
					SizeID: j.SizeID,
					Value:  j.Value,
				})
			}
			sizeList = append(sizeList, types.SizeList{
				ID:            i.ID,
				ProductID:     i.ProductID,
				Name:          i.Name,
				SizeValueList: sizeValueList,
			})
		}
		list = append(list, types.ListProductData{
			Id:            item.Id,
			CategoryID:    item.CategoryID,
			Name:          item.Name,
			Pic:           item.Pic,
			ProductSn:     item.ProductSn,
			Sale:          item.Sale,
			Price:         item.Price,
			SubTitle:      item.SubTitle,
			Description:   item.Description,
			OriginalPrice: item.OriginalPrice,
			Stock:         item.Stock,
			Unit:          item.Unit,
			SizeList:      sizeList,
			AttrValueIDs:  item.AttributeValueID,
		})
	}
	return &types.ListProductResp{
		Code:    200,
		Message: "查询列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
