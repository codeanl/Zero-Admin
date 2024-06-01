package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 优惠券列表
func (l *CouponListLogic) CouponList(in *sms.CouponListReq) (*sms.CouponListResp, error) {
	all, total, err := l.svcCtx.CouponModel.GetCouponList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.CouponListData
	for _, role := range all {
		list = append(list, &sms.CouponListData{
			Id:           int64(role.ID),
			Type:         role.Type,
			Name:         role.Name,
			Count:        role.Count,
			Amount:       role.Amount,
			PerLimit:     role.PerLimit,
			MinPoint:     role.MinPoint,
			StartTime:    role.StartTime,
			EndTime:      role.EndTime,
			UseType:      role.UseType,
			Note:         role.Note,
			UseCount:     role.UseCount,
			ReceiveCount: role.ReceiveCount,
			EnableTime:   role.EnableTime,
			Code:         role.Code,
		})
	}
	return &sms.CouponListResp{
		Total: total,
		List:  list,
	}, nil
}
