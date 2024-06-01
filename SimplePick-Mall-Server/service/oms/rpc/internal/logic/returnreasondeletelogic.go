package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnReasonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonDeleteLogic {
	return &ReturnReasonDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除退货原因
func (l *ReturnReasonDeleteLogic) ReturnReasonDelete(in *oms.ReturnReasonDeleteReq) (*oms.ReturnReasonDeleteResp, error) {
	err := l.svcCtx.ReturnReasonModel.DeleteReturnReasonByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &oms.ReturnReasonDeleteResp{}, nil
}
