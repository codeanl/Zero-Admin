package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyDeleteLogic {
	return &ReturnApplyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除退货
func (l *ReturnApplyDeleteLogic) ReturnApplyDelete(in *oms.ReturnApplyDeleteReq) (*oms.ReturnApplyDeleteResp, error) {
	err := l.svcCtx.ReturnApplyModel.DeleteReturnApplyByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &oms.ReturnApplyDeleteResp{}, nil
}
