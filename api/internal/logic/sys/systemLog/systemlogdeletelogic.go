package systemLog

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemLogDeleteLogic {
	return &SystemLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemLogDeleteLogic) SystemLogDelete(req *types.DeleteSysLogReq) (resp *types.DeleteSysLogResp, err error) {
	_, err = l.svcCtx.Sys.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除登录日志失败")
	}
	return &types.DeleteSysLogResp{
		Code:    200,
		Message: "删除登录日志成功",
	}, nil
}
