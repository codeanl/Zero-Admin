package order

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (*types.OrderInfoResp, error) {
	resp, _ := l.svcCtx.Oms.OrderInfo(l.ctx, &omsclient.OrderInfoReq{Id: req.ID})
	OrderInfo := types.ListOrderData{
		ID:                    resp.OrderInfo.ID,
		MemberId:              resp.OrderInfo.MemberId,
		PlaceId:               resp.OrderInfo.PlaceId,
		CouponId:              resp.OrderInfo.CouponId,
		OrderSn:               resp.OrderInfo.OrderSn,
		MemberUsername:        resp.OrderInfo.MemberUsername,
		TotalAmount:           resp.OrderInfo.TotalAmount,
		PayAmount:             resp.OrderInfo.PayAmount,
		FreightAmount:         resp.OrderInfo.FreightAmount,
		CouponAmount:          resp.OrderInfo.CouponAmount,
		PayType:               resp.OrderInfo.PayType,
		Status:                resp.OrderInfo.Status,
		OrderType:             resp.OrderInfo.OrderType,
		ReceiverProvince:      resp.OrderInfo.ReceiverProvince,
		ReceiverName:          resp.OrderInfo.ReceiverName,
		ReceiverPhone:         resp.OrderInfo.ReceiverPhone,
		ReceiverCity:          resp.OrderInfo.ReceiverCity,
		ReceiverRegion:        resp.OrderInfo.ReceiverRegion,
		ReceiverDetailAddress: resp.OrderInfo.ReceiverDetailAddress,
		Note:                  resp.OrderInfo.Note,
		ConfirmStatus:         resp.OrderInfo.ConfirmStatus,
		DeleteStatus:          resp.OrderInfo.DeleteStatus,
		PaymentTime:           resp.OrderInfo.PaymentTime,
		DeliveryTime:          resp.OrderInfo.DeliveryTime,
		ReceiveTime:           resp.OrderInfo.ReceiveTime,
		CommentTime:           resp.OrderInfo.CommentTime,
		CreateTime:            resp.OrderInfo.CreateTime,
	}
	var skuList []*types.SkuListData
	for _, i := range resp.Skus {
		idd := i.SkuID
		sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: idd})
		spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
		skuList = append(skuList, &types.SkuListData{
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
	place, _ := l.svcCtx.Sys.PlaceInfo(l.ctx, &sysclient.PlaceInfoReq{Id: resp.OrderInfo.PlaceId})
	PlaceInfo := types.PlaceInfoData{
		Id:        place.PlaceInfo.Id,
		Name:      place.PlaceInfo.Name,
		Place:     place.PlaceInfo.Place,
		Status:    place.PlaceInfo.Status,
		Pic:       place.PlaceInfo.Pic,
		Phone:     place.PlaceInfo.Phone,
		Principal: place.PlaceInfo.Principal,
	}
	sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: resp.Skus[0].SkuID})
	spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{ID: spu.ProductInfo.MerchantID})
	Merchant := types.MerchantInfo{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Principal: merchant.Principal,
		Phone:     merchant.Phone,
		Address:   merchant.Address,
		Pic:       merchant.Pic,
		UserID:    merchant.UserID,
	}
	data := types.OrderInfo{
		OrderInfo:    OrderInfo,
		SkuList:      skuList,
		PlaceInfo:    PlaceInfo,
		MerchantInfo: Merchant,
	}
	return &types.OrderInfoResp{
		Code:    200,
		Message: "success",
		Data:    data,
	}, nil
}
