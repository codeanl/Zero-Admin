package cart

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartListLogic) CartList(req *types.ListCartReq) (*types.ListCartResp, error) {
	resp, err := l.svcCtx.Oms.CartList(l.ctx, &oms.CartListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserID:   req.UserID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListCartData
	for _, item := range resp.List {
		var SkuData types.ListSkuData
		sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: item.SkuID})
		var data []map[string]string
		json.Unmarshal([]byte("["+sku.SkuInfo.Tag+"]"), &data)
		var space []types.SpaceInfo
		for _, i := range data {
			for key, value := range i {
				space = append(space, types.SpaceInfo{
					Name:      key,
					ValueName: value,
				})
			}
		}
		SkuData = types.ListSkuData{
			ID:          sku.SkuInfo.ID,
			ProductID:   sku.SkuInfo.ProductID,
			Name:        sku.SkuInfo.Name,
			Pic:         sku.SkuInfo.Pic,
			SkuSn:       sku.SkuInfo.SkuSn,
			Description: sku.SkuInfo.Description,
			Price:       sku.SkuInfo.Price,
			Stock:       sku.SkuInfo.Stock,
			Tag:         sku.SkuInfo.Tag,
			TagText:     sku.SkuInfo.Tag,
			Space:       space,
		}
		spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
		merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{ID: spu.ProductInfo.MerchantID})
		Merchant := types.MerchantData{
			ID:        merchant.ID,
			Name:      merchant.Name,
			Principal: merchant.Principal,
			Phone:     merchant.Phone,
			Address:   merchant.Address,
			Pic:       merchant.Pic,
			UserID:    merchant.UserID,
		}
		listUserData := types.ListCartData{
			ID:           item.ID,
			UserID:       item.UserID,
			SkuID:        item.SkuID,
			Count:        item.Count,
			ListSkuData:  SkuData,
			MerchantData: Merchant,
		}
		list = append(list, listUserData)
	}

	return &types.ListCartResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
