package logic

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq, ip string) (*types.LoginResp, error) {
	resp, err := l.svcCtx.Sys.UserLogin(l.ctx, &sysclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("登录失败")
	}
	//登录日志
	_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
		UserID: resp.UserID,
		Status: "online",
		IP:     ip,
	})
	return &types.LoginResp{
		Code:    200,
		Data:    resp.Token,
		Message: "success",
	}, nil
}
