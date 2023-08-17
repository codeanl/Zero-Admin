package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加订单
func (l *OrderAddLogic) OrderAdd(in *oms.OrderAddReq) (*oms.OrderAddResp, error) {
	role := &model.Order{
		PlaceId:               in.PlaceId,
		MemberId:              in.MemberId,
		OrderSn:               in.OrderSn,
		MemberUsername:        in.MemberUsername,
		TotalAmount:           in.TotalAmount,
		PayAmount:             in.PayAmount,
		FreightAmount:         in.FreightAmount,
		CouponAmount:          in.CouponAmount,
		PayType:               in.PayType,
		Status:                in.Status,
		OrderType:             in.OrderType,
		ReceiverName:          in.ReceiverName,
		ReceiverPhone:         in.ReceiverPhone,
		ReceiverProvince:      in.ReceiverProvince,
		ReceiverCity:          in.ReceiverCity,
		ReceiverRegion:        in.ReceiverRegion,
		ReceiverDetailAddress: in.ReceiverDetailAddress,
		Note:                  in.Note,
		ConfirmStatus:         in.ConfirmStatus,
		DeleteStatus:          in.DeleteStatus,
		PaymentTime:           in.PaymentTime,
	}
	order, err := l.svcCtx.OrderModel.AddOrder(role)
	if err != nil {
		//return nil, errors.New("添加失败")
		return nil, err
	}
	for _, i := range in.SkuIDs {
		l.svcCtx.OrderSkuModel.AddOrderSku(&model.OrderSku{
			OrderID: int64(order.ID),
			SkuID:   i,
		})
	}
	return &oms.OrderAddResp{}, nil
}
