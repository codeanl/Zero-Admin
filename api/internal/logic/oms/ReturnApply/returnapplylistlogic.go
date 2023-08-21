package ReturnApply

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"log"

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
		//return nil, errorx.NewDefaultError("查询失败")
		return nil, err
	}
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
		for _, i := range order.SkuIDs {
			sku, _ := l.svcCtx.Pms.SkuInfo(l.ctx, &pmsclient.SkuInfoReq{ID: i})
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
		user, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: item.UserID})
		log.Print(user)
		User := types.UserData{
			ID:       user.UserInfo.ID,
			Username: user.UserInfo.Username,
			NickName: user.UserInfo.Nickname,
			Phone:    user.UserInfo.Phone,
			Gander:   user.UserInfo.Gender,
			Avatar:   user.UserInfo.Avatar,
			Email:    user.UserInfo.Email,
			Status:   user.UserInfo.Status,
			CreateAt: user.UserInfo.CreateAt,
			UpdateAt: user.UserInfo.UpdateAt,
			CreateBy: user.UserInfo.CreateBy,
			UpdateBy: user.UserInfo.UpdateBy,
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
		list = append(list, listUserData)
	}
	return &types.ListReturnApplyResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
