package systemLog

import (
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
		return &types.DeleteSysLogResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteSysLogResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
