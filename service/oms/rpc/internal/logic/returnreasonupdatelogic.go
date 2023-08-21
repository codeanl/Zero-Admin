package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnReasonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonUpdateLogic {
	return &ReturnReasonUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新退货原因
func (l *ReturnReasonUpdateLogic) ReturnReasonUpdate(in *oms.ReturnReasonUpdateReq) (*oms.ReturnReasonUpdateResp, error) {
	err := l.svcCtx.ReturnReasonModel.UpdateReturnReason(in.ID, &model.ReturnReason{
		Name:   in.Name,
		Status: in.Status,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &oms.ReturnReasonUpdateResp{}, nil
}
