package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnReasonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonAddLogic {
	return &ReturnReasonAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加退货原因
func (l *ReturnReasonAddLogic) ReturnReasonAdd(in *oms.ReturnReasonAddReq) (*oms.ReturnReasonAddResp, error) {
	role := &model.ReturnReason{
		Name:   in.Name,
		Status: in.Status,
	}
	_, err := l.svcCtx.ReturnReasonModel.AddReturnReason(role)
	if err != nil {
		//return nil, errors.New("添加失败")
		return nil, err
	}
	return &oms.ReturnReasonAddResp{}, nil
}
