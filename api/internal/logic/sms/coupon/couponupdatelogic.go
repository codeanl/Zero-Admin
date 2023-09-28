package coupon

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponUpdateLogic {
	return &CouponUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponUpdateLogic) CouponUpdate(req *types.UpdateCouponReq) (resp *types.UpdateCouponResp, err error) {
	_, err = l.svcCtx.Sms.CouponUpdate(l.ctx, &smsclient.CouponUpdateReq{
		Id:         req.Id,
		Type:       req.Type,
		Name:       req.Name,
		Count:      req.Count,
		Amount:     req.Amount,
		PerLimit:   req.PerLimit,
		MinPoint:   req.MinPoint,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		UseType:    req.UseType,
		Note:       req.Note,
		EnableTime: req.EnableTime,
		Code:       req.Code,
	})
	if err != nil {
		return &types.UpdateCouponResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateCouponResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
