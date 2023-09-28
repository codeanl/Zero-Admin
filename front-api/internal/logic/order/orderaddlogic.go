package order

import (
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req *types.AddOrderReq) (*types.AddOrderResp, error) {
	//var Skus []*omsclient.Skus
	//for _, i := range req.Skus {
	//	Skus = append(Skus, &omsclient.Skus{
	//		Count: i.Count,
	//		SkuID: i.SkuID,
	//	})
	//	l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i.SkuID})
	//}
	//order, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
	//	PlaceId:               req.PlaceId,
	//	MemberId:              req.MemberId,
	//	OrderSn:               req.OrderSn,
	//	MemberUsername:        req.MemberUsername,
	//	TotalAmount:           req.TotalAmount,
	//	PayAmount:             req.PayAmount,
	//	FreightAmount:         req.FreightAmount,
	//	CouponAmount:          req.CouponAmount,
	//	PayType:               req.PayType,
	//	Status:                req.Status,
	//	OrderType:             req.OrderType,
	//	ReceiverName:          req.ReceiverName,
	//	ReceiverPhone:         req.ReceiverPhone,
	//	ReceiverProvince:      req.ReceiverProvince,
	//	ReceiverCity:          req.ReceiverCity,
	//	ReceiverRegion:        req.ReceiverRegion,
	//	ReceiverDetailAddress: req.ReceiverDetailAddress,
	//	Note:                  req.Note,
	//	ConfirmStatus:         "0", //确认收货状态：0->未确认
	//	DeleteStatus:          "0", // 删除状态：0->未删除
	//	PaymentTime:           req.PaymentTime,
	//	Skus:                  Skus,
	//})
	//data := types.Data{
	//	OrderID: order.OrderID,
	//}
	//if err != nil {
	//	return nil, err
	//}
	type SkuList struct {
		Sku        []*omsclient.Skus
		MerchantID int64
	}
	var data []SkuList
	for _, i := range req.Skus {
		skuInfo, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i.SkuID})
		aaa := &omsclient.Skus{
			Count: i.Count,
			SkuID: i.SkuID,
		}
		if len(data) == 0 {
			var sku []*omsclient.Skus
			sku = append(sku, aaa)
			skuss := SkuList{
				MerchantID: skuInfo.SkuInfo.MerchantID,
				Sku:        sku,
			}
			data = append(data, skuss)
		} else {
			for index, j := range data {
				if skuInfo.SkuInfo.MerchantID == j.MerchantID {
					data[index].Sku = append(data[index].Sku, aaa)
				} else {
					var sku []*omsclient.Skus
					sku = append(sku, aaa)
					skuss := SkuList{
						MerchantID: skuInfo.SkuInfo.MerchantID,
						Sku:        sku,
					}
					data = append(data, skuss)
				}
			}
		}
	}
	log.Print(data)
	var orderIDS []int64
	for _, i := range data {
		var Skus []*omsclient.Skus
		for _, j := range i.Sku {
			Skus = append(Skus, &omsclient.Skus{
				Count: j.Count,
				SkuID: j.SkuID,
			})
		}
		order, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
			PlaceId:               req.PlaceId,
			MemberId:              req.MemberId,
			OrderSn:               req.OrderSn,
			MemberUsername:        req.MemberUsername,
			TotalAmount:           req.TotalAmount,
			PayAmount:             req.PayAmount,
			FreightAmount:         req.FreightAmount,
			CouponAmount:          req.CouponAmount,
			PayType:               req.PayType,
			Status:                req.Status,
			OrderType:             req.OrderType,
			ReceiverName:          req.ReceiverName,
			ReceiverPhone:         req.ReceiverPhone,
			ReceiverProvince:      req.ReceiverProvince,
			ReceiverCity:          req.ReceiverCity,
			ReceiverRegion:        req.ReceiverRegion,
			ReceiverDetailAddress: req.ReceiverDetailAddress,
			Note:                  req.Note,
			ConfirmStatus:         "0", //确认收货状态：0->未确认
			DeleteStatus:          "0", // 删除状态：0->未删除
			PaymentTime:           req.PaymentTime,
			Skus:                  Skus,
		})
		if err != nil {
			return nil, err
		}
		orderIDS = append(orderIDS, order.OrderID)
	}
	datas := types.Data{
		OrderID: orderIDS,
	}
	return &types.AddOrderResp{
		Code:    200,
		Message: "添加成功",
		Data:    datas,
	}, nil
}
