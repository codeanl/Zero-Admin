package order

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		ID:                    req.ID,
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
		ConfirmStatus:         req.ConfirmStatus,
		DeleteStatus:          req.DeleteStatus,
		PaymentTime:           req.PaymentTime,
		DeliveryTime:          req.DeliveryTime,
		ReceiveTime:           req.ReceiveTime,
		CommentTime:           req.CommentTime,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	order, err := l.svcCtx.Oms.OrderInfo(l.ctx, &omsclient.OrderInfoReq{Id: req.ID})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败,查询用户的时候出错")
	}
	if req.Status == "1" {
		for _, i := range order.Skus {
			sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i.SkuID})
			l.svcCtx.Pms.SkuUpdate(l.ctx, &pmsclient.SkuUpdateReq{
				ID:    i.SkuID,
				Sale:  sku.SkuInfo.Sale + 1,
				Stock: sku.SkuInfo.Stock - 1,
			})
		}
	}
	return &types.UpdateOrderResp{
		Code:    200,
		Message: "更新用户成功",
	}, nil
}
