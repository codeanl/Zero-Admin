package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单
func (l *OrderUpdateLogic) OrderUpdate(in *oms.OrderUpdateReq) (*oms.OrderUpdateResp, error) {
	err := l.svcCtx.OrderModel.UpdateOrder(in.ID, &model.Order{
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
		DeliveryTime:          in.DeliveryTime,
		ReceiveTime:           in.ReceiveTime,
		CommentTime:           in.CommentTime,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &oms.OrderUpdateResp{}, nil

}
