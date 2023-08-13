package order

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.ListOrderReq) (*types.ListOrderResp, error) {
	resp, err := l.svcCtx.Oms.OrderList(l.ctx, &omsclient.OrderListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		//return nil, errorx.NewDefaultError("查询失败")
		return nil, err
	}
	var list []*types.ListOrderData
	for _, item := range resp.List {
		listUserData := types.ListOrderData{
			Id:                    item.ID,
			PlaceId:               item.PlaceId,
			MemberId:              item.MemberId,
			OrderSn:               item.OrderSn,
			MemberUsername:        item.MemberUsername,
			TotalAmount:           item.TotalAmount,
			PayAmount:             item.PayAmount,
			FreightAmount:         item.FreightAmount,
			CouponAmount:          item.CouponAmount,
			DiscountAmount:        item.DiscountAmount,
			PayType:               item.PayType,
			Status:                item.Status,
			OrderType:             item.OrderType,
			AutoConfirmDay:        item.AutoConfirmDay,
			ReceiverProvince:      item.ReceiverProvince,
			ReceiverCity:          item.ReceiverCity,
			ReceiverRegion:        item.ReceiverRegion,
			ReceiverDetailAddress: item.ReceiverDetailAddress,
			Note:                  item.Note,
			ConfirmStatus:         item.ConfirmStatus,
			DeleteStatus:          item.DeleteStatus,
			DeliveryTime:          item.DeliveryTime,
			ReceiveTime:           item.ReceiveTime,
			CommentTime:           item.CommentTime,
		}
		list = append(list, &listUserData)
	}
	return &types.ListOrderResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
