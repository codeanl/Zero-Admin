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
		Current:        req.Current,
		PageSize:       req.PageSize,
		OrderSn:        req.OrderSn,
		Status:         req.Status,
		MemberUsername: req.MemberUsername,
	})
	if err != nil {
		//return nil, errorx.NewDefaultError("查询失败")
		return nil, err
	}
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	place, _ := l.svcCtx.Sys.PlaceInfo(l.ctx, &sysclient.PlaceInfoReq{UserID: userInfo.UserInfo.ID})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
	isZTD := false
	for _, ii := range userInfo.Roles {
		if ii == "自提点管理员" {
			isZTD = true
		}
	}
	isSJ := false
	for _, ii := range userInfo.Roles {
		if ii == "商家" {
			isSJ = true
		}
	}
	log.Print(merchant.ID)
	var list []types.ListOrderData
	for _, item := range resp.List {
		sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: item.Skus[0].SkuID})
		if isZTD { //
			if item.PlaceId == place.PlaceInfo.Id {
				listUserData := types.ListOrderData{
					ID:                    item.ID,
					PlaceId:               item.PlaceId,
					MemberId:              item.MemberId,
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
			} //
		} else if isSJ {
			if merchant.ID == sku.SkuInfo.MerchantID {
				listUserData := types.ListOrderData{
					ID:                    item.ID,
					PlaceId:               item.PlaceId,
					MemberId:              item.MemberId,
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
		} else {
			listUserData := types.ListOrderData{
				ID:                    item.ID,
				PlaceId:               item.PlaceId,
				MemberId:              item.MemberId,
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

	}
	return &types.ListOrderResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
