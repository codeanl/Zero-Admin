package order

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req *types.UpdateOrderReq) (resp *types.UpdateOrderResp, err error) {
	_, err = l.svcCtx.Oms.OrderUpdate(l.ctx, &omsclient.OrderUpdateReq{
		ID:                    req.Id,
		PlaceId:               req.PlaceId,
		MemberId:              req.MemberId,
		OrderSn:               req.OrderSn,
		MemberUsername:        req.MemberUsername,
		TotalAmount:           req.TotalAmount,
		PayAmount:             req.PayAmount,
		FreightAmount:         req.FreightAmount,
		CouponAmount:          req.CouponAmount,
		DiscountAmount:        req.DiscountAmount,
		PayType:               req.PayType,
		Status:                req.Status,
		OrderType:             req.OrderType,
		AutoConfirmDay:        req.AutoConfirmDay,
		ReceiverProvince:      req.ReceiverProvince,
		ReceiverCity:          req.ReceiverCity,
		ReceiverRegion:        req.ReceiverRegion,
		ReceiverDetailAddress: req.ReceiverDetailAddress,
		Note:                  req.Note,
		ConfirmStatus:         req.ConfirmStatus,
		DeleteStatus:          req.DeleteStatus,
		DeliveryTime:          req.DeliveryTime,
		ReceiveTime:           req.ReceiveTime,
		CommentTime:           req.CommentTime,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateOrderResp{
		Code:    200,
		Message: "更新用户成功",
	}, nil
}
