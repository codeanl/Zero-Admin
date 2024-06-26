package product

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
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

	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.UserInfoReq{Id: id})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if strings.Contains("商家", ii) {
			isSJ = true
		}
	}
	if isSJ == true {
		merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: id})
		req.MerchantID = merchant.ID
	}
	resp, _ := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		Name:       req.Name,
		CategoryID: req.CategoryId,
		MerchantID: req.MerchantID,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		SearchType: req.SearchType,
	})
	//if err != nil {
	//	return &types.ListProductResp{
	//		Code:    400,
	//		Message: "查询失败",
	//	}, nil
	//}

	var list []types.ListProductData
	for _, item := range resp.List {
		//
		skuList, _ := l.svcCtx.Pms.SkuList(l.ctx, &pmsclient.SkuListReq{ProductID: item.Id})
		var sale int64 = 0
		for _, i := range skuList.List {
			sale = sale + i.Sale
		}
		//_, _ = l.svcCtx.Pms.ProductUpdate(l.ctx, &pmsclient.ProductUpdateReq{
		//	Id:   item.Id,
		//	Sale: sale,
		//})

		listProduct := types.ListProductData{
			Id:                  item.Id,
			CategoryID:          item.CategoryID,
			Name:                item.Name,
			Pic:                 item.Pic,
			ProductSn:           item.ProductSn,
			Price:               item.Price,
			Desc:                item.Desc,
			OriginalPrice:       item.OriginalPrice,
			Unit:                item.Unit,
			AttributeCategoryID: item.AttributeCategoryID,
			MerchantID:          item.MerchantID,
			Sale:                sale,
		}
		list = append(list, listProduct)
	}
	return &types.ListProductResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
