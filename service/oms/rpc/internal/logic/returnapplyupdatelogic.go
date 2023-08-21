package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyUpdateLogic {
	return &ReturnApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新退货
func (l *ReturnApplyUpdateLogic) ReturnApplyUpdate(in *oms.ReturnApplyUpdateReq) (*oms.ReturnApplyUpdateResp, error) {
	err := l.svcCtx.ReturnApplyModel.UpdateReturnApply(in.ID, &model.ReturnApply{
		UserID:         in.UserID,
		OrderID:        in.OrderID,
		ReturnReasonID: in.ReturnReasonID,
		Status:         in.Status,
		Description:    in.Description,
		ProofPics:      in.ProofPics,
		ReturnAmount:   in.ReturnAmount,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &oms.ReturnApplyUpdateResp{}, nil
}
