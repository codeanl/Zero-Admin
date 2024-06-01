package order

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"strings"
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
	//
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.UserInfoReq{Id: id})
	place, _ := l.svcCtx.Pms.PlaceInfo(l.ctx, &pmsclient.PlaceInfoReq{UserID: userInfo.UserInfo.Id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.Id})
	isZTD := false
	for _, ii := range userInfo.Roles {
		if strings.Contains("自提点管理员", ii) {
			isZTD = true
		}
	}
	log.Print(isZTD)
	if isZTD {
		req.PlaceID = place.PlaceInfo.Id
	}
	isSJ := false
	for _, ii := range userInfo.Roles {
		if strings.Contains("商家", ii) {
			isSJ = true
		}
	}
	if isSJ {
		req.MerchantID = merchant.ID
	}
	//
	resp, err := l.svcCtx.Oms.OrderList(l.ctx, &omsclient.OrderListReq{
		Current:        req.Current,
		PageSize:       req.PageSize,
		OrderSn:        req.OrderSn,
		Status:         req.Status,
		MemberUsername: req.MemberUsername,
		PlaceID:        req.PlaceID,
		MerchantID:     req.MerchantID,
	})
	if err != nil {
		return &types.ListOrderResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []types.ListOrderData
	for _, item := range resp.List {
		listUserData := types.ListOrderData{
			ID:                    item.ID,
			PlaceId:               item.PlaceId,
			MemberId:              item.MemberId,
			MerchantID:            item.MerchantID,
			OrderSn:               item.OrderSn,
			MemberUsername:        item.MemberUsername,
			TotalAmount:           item.TotalAmount,
			PayAmount:             item.PayAmount,
			FreightAmount:         item.FreightAmount,
			CouponAmount:          item.CouponAmount,
			PayType:               item.PayType,
			Status:                item.Status,
			OrderType:             item.OrderType,
			ReceiverName:          item.ReceiverName,
			ReceiverPhone:         item.ReceiverPhone,
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
		list = append(list, listUserData)
	}
	return &types.ListOrderResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
