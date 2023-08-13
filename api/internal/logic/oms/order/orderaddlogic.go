package order

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

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

func (l *OrderAddLogic) OrderAdd(req *types.AddOrderReq) (resp *types.AddOrderResp, err error) {
	_, err = l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
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
	})
	if err != nil {
		return nil, err
	}
	return &types.AddOrderResp{
		Code:    200,
		Message: "添加成功",
	}, nil
	return
}
