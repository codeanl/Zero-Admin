package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除日志
func (l *SysLogDeleteLogic) SysLogDelete(in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	err := l.svcCtx.LogModel.DeleteLogByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除系统日志失败")
	}
	return &sys.SysLogDeleteResp{}, nil
}
