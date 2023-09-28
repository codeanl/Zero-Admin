package user

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestartPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartPasswordLogic {
	return &RestartPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestartPasswordLogic) RestartPassword(req *types.RestartPasswordReq) (resp *types.RestartPasswordResp, err error) {
	_, err = l.svcCtx.Sys.RestartPassword(l.ctx, &sysclient.RestartPasswordReq{ID: req.ID})
	return &types.RestartPasswordResp{
		Code:    200,
		Message: "重置密码成功：123456",
	}, nil
}
