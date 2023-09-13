package order

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
	var Skus []*omsclient.Skus
	for _, i := range req.Skus {
		Skus = append(Skus, &omsclient.Skus{
			Count: i.Count,
			SkuID: i.SkuID,
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
	data := types.Data{
		OrderID: order.OrderID,
	}
	if err != nil {
		return nil, err
	}
	return &types.AddOrderResp{
		Code:    200,
		Message: "添加成功",
		Data:    data,
	}, nil
}
