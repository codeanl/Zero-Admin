package order

import (
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.ListOrderReq) (*types.ListOrderResp, error) {
	resp, err := l.svcCtx.Oms.OrderList(l.ctx, &omsclient.OrderListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
		UserID:   req.UserID,
	})
	if err != nil {
		//return nil, errorx.NewDefaultError("查询失败")
		return nil, err
	}
	var list []types.ListOrderData
	for _, item := range resp.List {
		var skuList []types.Sku
		for _, i := range item.Skus {
			sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i.SkuID})
			spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
			skuList = append(skuList, types.Sku{
				ID:          sku.SkuInfo.ID,
				ProductID:   sku.SkuInfo.ProductID,
				ProductName: spu.ProductInfo.Name,
				Name:        sku.SkuInfo.Name,
				Pic:         sku.SkuInfo.Pic,
				SkuSn:       sku.SkuInfo.SkuSn,
				Description: sku.SkuInfo.Description,
				Price:       sku.SkuInfo.Price,
				Stock:       sku.SkuInfo.Stock,
				Tag:         sku.SkuInfo.Tag,
				Count:       i.Count,
			})
		}
		sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: item.Skus[0].SkuID})
		spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
		merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{ID: spu.ProductInfo.MerchantID})
		merchantInfo := types.MerchantInfo{
			ID:        merchant.ID,
			Name:      merchant.Name,
			Principal: merchant.Principal,
			Phone:     merchant.Phone,
			Address:   merchant.Address,
			Pic:       merchant.Pic,
			UserID:    merchant.UserID,
		}
		list = append(list, types.ListOrderData{
			ID:                    item.ID,
			PlaceId:               item.PlaceId,
			MemberId:              item.MemberId,
			OrderSn:               item.OrderSn,
			MemberUsername:        item.MemberUsername,
			TotalAmount:           item.TotalAmount,
			PayAmount:             item.PayAmount,
			FreightAmount:         item.FreightAmount,
			CouponAmount:          item.CouponAmount,
			PayType:               item.PayType,
			Status:                item.Status,
			OrderType:             item.OrderType,
			ReceiverName:          item.ReceiverName,
			ReceiverPhone:         item.ReceiverPhone,
			ReceiverProvince:      item.ReceiverProvince,
			ReceiverCity:          item.ReceiverCity,
			ReceiverRegion:        item.ReceiverRegion,
			ReceiverDetailAddress: item.ReceiverDetailAddress,
			Note:                  item.Note,
			ConfirmStatus:         item.ConfirmStatus,
			DeleteStatus:          item.DeleteStatus,
			DeliveryTime:          item.DeliveryTime,
			CreateTime:            item.CreateTime,
			ReceiveTime:           item.ReceiveTime,
			CommentTime:           item.CommentTime,
			SkuList:               skuList,
			MerchantInfo:          merchantInfo,
		})
	}
	return &types.ListOrderResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
