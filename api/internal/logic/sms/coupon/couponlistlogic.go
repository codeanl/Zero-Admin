package coupon

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponListLogic) CouponList(req *types.ListCouponReq) (*types.ListCouponResp, error) {
	resp, err := l.svcCtx.Sms.CouponList(l.ctx, &smsclient.CouponListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Type:     req.Type,
		UseType:  req.UseType,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListCouponData
	for _, item := range resp.List {
		listUserData := types.ListCouponData{
			Id:         item.Id,
			Type:       item.Type,
			Name:       item.Name,
			Count:      item.Count,
			Amount:     item.Amount,
			PerLimit:   item.PerLimit,
			MinPoint:   item.MinPoint,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			UseType:    item.UseType,
			Note:       item.Note,
			EnableTime: item.EnableTime,
			Code:       item.Code,
		}
		list = append(list, &listUserData)
	}

	return &types.ListCouponResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
