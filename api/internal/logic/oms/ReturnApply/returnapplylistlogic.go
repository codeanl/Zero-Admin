package ReturnApply

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyListLogic {
	return &ReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyListLogic) ReturnApplyList(req *types.ListReturnApplyReq) (*types.ListReturnApplyResp, error) {
	resp, err := l.svcCtx.Oms.ReturnApplyList(l.ctx, &omsclient.ReturnApplyListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
	})
	if err != nil {

		return &types.ListReturnApplyResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	//
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
	//
	var list []types.ListReturnApplyData
	for _, item := range resp.List {
		order, _ := l.svcCtx.Oms.OrderInfo(l.ctx, &omsclient.OrderInfoReq{Id: item.OrderID})
		OrderInfo := types.OrderData{
			ID:                    order.OrderInfo.ID,
			MemberId:              order.OrderInfo.MemberId,
			PlaceId:               order.OrderInfo.PlaceId,
			CouponId:              order.OrderInfo.CouponId,
			OrderSn:               order.OrderInfo.OrderSn,
			MemberUsername:        order.OrderInfo.MemberUsername,
			TotalAmount:           order.OrderInfo.TotalAmount,
			PayAmount:             order.OrderInfo.PayAmount,
			FreightAmount:         order.OrderInfo.FreightAmount,
			CouponAmount:          order.OrderInfo.CouponAmount,
			PayType:               order.OrderInfo.PayType,
			Status:                order.OrderInfo.Status,
			OrderType:             order.OrderInfo.OrderType,
			ReceiverProvince:      order.OrderInfo.ReceiverProvince,
			ReceiverName:          order.OrderInfo.ReceiverName,
			ReceiverPhone:         order.OrderInfo.ReceiverPhone,
			ReceiverCity:          order.OrderInfo.ReceiverCity,
			ReceiverRegion:        order.OrderInfo.ReceiverRegion,
			ReceiverDetailAddress: order.OrderInfo.ReceiverDetailAddress,
			Note:                  order.OrderInfo.Note,
			ConfirmStatus:         order.OrderInfo.ConfirmStatus,
			DeleteStatus:          order.OrderInfo.DeleteStatus,
			PaymentTime:           order.OrderInfo.PaymentTime,
			DeliveryTime:          order.OrderInfo.DeliveryTime,
			ReceiveTime:           order.OrderInfo.ReceiveTime,
			CommentTime:           order.OrderInfo.CommentTime,
		}
		var skuList []types.SkuData
		for _, i := range order.Skus {
			sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i.SkuID})
			spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
			skuList = append(skuList, types.SkuData{
				ID:          sku.SkuInfo.ID,
				ProductID:   sku.SkuInfo.ProductID,
				ProductName: spu.ProductInfo.Name,
				Name:        sku.SkuInfo.Name,
				Pic:         sku.SkuInfo.Pic,
				SkuSn:       sku.SkuInfo.SkuSn,
				Description: sku.SkuInfo.Description,
				Price:       sku.SkuInfo.Price,
				Stock:       sku.SkuInfo.Stock,
				Tag:         sku.SkuInfo.Tag,
			})
		}
		Order := types.Order{
			OrderInfo: OrderInfo,
			SkuList:   skuList,
		}
		user, _ := l.svcCtx.Ums.MemberInfo(l.ctx, &umsclient.MemberInfoReq{Id: item.UserID})
		User := types.UserData{
			ID:       user.Id,
			Username: user.Username,
			NickName: user.Nickname,
			Phone:    user.Phone,
			Gander:   user.Gender,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Status:   user.Status,
		}
		listUserData := types.ListReturnApplyData{
			ID:               item.ID,
			UserID:           item.UserID,
			OrderID:          item.OrderID,
			ReturnReasonID:   item.ReturnReasonID,
			ReturnReasonName: item.ReturnReasonName,
			Status:           item.Status,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			ReturnAmount:     item.ReturnAmount,
			Order:            Order,
			User:             User,
		}
		//
		if isZTD {
			if order.OrderInfo.PlaceId == place.PlaceInfo.Id {
				list = append(list, listUserData)
			}
		} else if isSJ {
			sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: order.Skus[0].SkuID})
			spu, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: sku.SkuInfo.ProductID})
			if spu.ProductInfo.MerchantID == merchant.ID {
				list = append(list, listUserData)
			}
		} else {
			list = append(list, listUserData)
		}
		//
		//list = append(list, listUserData)
	}
	return &types.ListReturnApplyResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
