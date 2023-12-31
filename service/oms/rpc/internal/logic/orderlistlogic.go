package logic

import (
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 订单列表
func (l *OrderListLogic) OrderList(in *oms.OrderListReq) (*oms.OrderListResp, error) {
	all, total, err := l.svcCtx.OrderModel.GetOrderList(in)
	if err != nil {
		return nil, err
	}
	var list []*oms.OrderListData
	for _, role := range all {
		var skus []*oms.Skus
		sku, _ := l.svcCtx.OrderSkuModel.GetOrderSkuByOrderID(int64(role.ID))
		for _, i := range sku {
			skus = append(skus, &oms.Skus{
				SkuID: i.SkuID,
				Count: i.Count,
			})
		}
		list = append(list, &oms.OrderListData{
			ID:                    int64(role.ID),
			PlaceId:               role.PlaceId,
			MemberId:              role.MemberId,
			OrderSn:               role.OrderSn,
			MemberUsername:        role.MemberUsername,
			TotalAmount:           role.TotalAmount,
			PayAmount:             role.PayAmount,
			FreightAmount:         role.FreightAmount,
			CouponAmount:          role.CouponAmount,
			PayType:               role.PayType,
			Status:                role.Status,
			OrderType:             role.OrderType,
			ReceiverName:          role.ReceiverName,
			ReceiverPhone:         role.ReceiverPhone,
			ReceiverProvince:      role.ReceiverProvince,
			ReceiverCity:          role.ReceiverCity,
			ReceiverRegion:        role.ReceiverRegion,
			ReceiverDetailAddress: role.ReceiverDetailAddress,
			Note:                  role.Note,
			ConfirmStatus:         role.ConfirmStatus,
			CreateTime:            role.CreatedAt.Format("2006-01-02 15:04:05"),
			DeleteStatus:          role.DeleteStatus,
			PaymentTime:           role.PaymentTime,
			DeliveryTime:          role.DeliveryTime,
			ReceiveTime:           role.ReceiveTime,
			CommentTime:           role.CommentTime,
			Skus:                  skus,
		})
	}
	return &oms.OrderListResp{
		Total: total,
		List:  list,
	}, nil
}
