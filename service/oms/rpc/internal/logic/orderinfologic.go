package logic

import (
	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单详情
func (l *OrderInfoLogic) OrderInfo(in *oms.OrderInfoReq) (*oms.OrderInfoResp, error) {
	order, err := l.svcCtx.OrderModel.GetOrderById(in.Id)
	if err != nil {
		return nil, err
	}
	info := &oms.OrderListData{
		ID:                    int64(order.ID),
		PlaceId:               order.PlaceId,
		MemberId:              order.MemberId,
		CouponId:              order.CouponId,
		OrderSn:               order.OrderSn,
		MemberUsername:        order.MemberUsername,
		TotalAmount:           order.TotalAmount,
		PayAmount:             order.PayAmount,
		FreightAmount:         order.FreightAmount,
		CouponAmount:          order.CouponAmount,
		PayType:               order.PayType,
		Status:                order.Status,
		OrderType:             order.OrderType,
		ReceiverName:          order.ReceiverName,
		ReceiverPhone:         order.ReceiverPhone,
		ReceiverProvince:      order.ReceiverProvince,
		ReceiverCity:          order.ReceiverCity,
		ReceiverRegion:        order.ReceiverRegion,
		ReceiverDetailAddress: order.ReceiverDetailAddress,
		Note:                  order.Note,
		ConfirmStatus:         order.ConfirmStatus,
		DeleteStatus:          order.DeleteStatus,
		PaymentTime:           order.PaymentTime,
		DeliveryTime:          order.DeliveryTime,
		ReceiveTime:           order.ReceiveTime,
		CommentTime:           order.CommentTime,
	}
	log.Print(info)
	var SkuIDs []int64
	sku, err := l.svcCtx.OrderSkuModel.GetOrderSkuByOrderID(in.Id)
	if err != nil {
		return nil, err
	}
	for _, i := range sku {
		SkuIDs = append(SkuIDs, i.SkuID)
	}
	return &oms.OrderInfoResp{
		OrderInfo: info,
		SkuIDs:    SkuIDs,
	}, nil
}
