package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加优惠券
func (l *CouponAddLogic) CouponAdd(in *sms.CouponAddReq) (*sms.CouponAddResp, error) {
	role := &model.Coupon{
		Type:         in.Type,
		Name:         in.Name,
		Count:        in.Count,
		Amount:       float64(in.Amount),
		PerLimit:     in.PerLimit,
		MinPoint:     float64(in.MinPoint),
		StartTime:    in.StartTime,
		EndTime:      in.EndTime,
		UseType:      in.UseType,
		Note:         in.Note,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   in.EnableTime,
		Code:         in.Code,
	}
	err := l.svcCtx.CouponModel.AddCoupon(role)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &sms.CouponAddResp{}, nil
}
