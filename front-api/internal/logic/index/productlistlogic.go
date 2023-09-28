package index

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		CategoryID: req.CategoryId,
		Name:       req.Name,
		MerchantID: req.MerchantID,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		SearchType: req.SearchType,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListProductData
	for _, item := range resp.List {
		skuList, _ := l.svcCtx.Pms.SkuList(l.ctx, &pmsclient.SkuListReq{ProductID: item.Id})
		var sale int64 = 0
		for _, i := range skuList.List {
			sale = sale + i.Sale
		}
		_, _ = l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
			Id:   item.Id,
			Sale: sale,
		})
		list = append(list, types.ListProductData{
			Id:            item.Id,
			CategoryID:    item.CategoryID,
			Name:          item.Name,
			Pic:           item.Pic,
			ProductSn:     item.ProductSn,
			Price:         item.Price,
			Desc:          item.Desc,
			OriginalPrice: item.OriginalPrice,
			Unit:          item.Unit,
		})
	}
	return &types.ListProductResp{
		Code:    200,
		Message: "查询列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
