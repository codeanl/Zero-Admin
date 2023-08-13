package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

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
		DiscountAmount:        in.DiscountAmount,
		PayType:               in.PayType,
		Status:                in.Status,
		OrderType:             in.OrderType,
		AutoConfirmDay:        in.AutoConfirmDay,
		ReceiverProvince:      in.ReceiverProvince,
		ReceiverCity:          in.ReceiverCity,
		ReceiverRegion:        in.ReceiverRegion,
		ReceiverDetailAddress: in.ReceiverDetailAddress,
		Note:                  in.Note,
		ConfirmStatus:         in.ConfirmStatus,
		DeleteStatus:          in.DeleteStatus,
		PaymentTime:           in.PaymentTime,
		DeliveryTime:          in.DeliveryTime,
		ReceiveTime:           in.ReceiveTime,
		CommentTime:           in.CommentTime,
	}
	err := l.svcCtx.OrderModel.AddOrder(role)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}

	return &oms.OrderAddResp{}, nil

}
