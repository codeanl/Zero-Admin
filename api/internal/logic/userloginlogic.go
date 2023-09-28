package logic

import (
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
		return &types.LoginResp{
			Code:    400,
			Message: "账号或密码有误，请重新输入",
		}, nil
	}
	if resp.Status == "0" {
		return &types.LoginResp{
			Code:    400,
			Message: "账号已锁定，请联系超级管理员解锁。",
		}, nil
	}
	//登录日志
	_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
		UserID: resp.UserID,
		Status: "1",
		IP:     ip,
	})
	return &types.LoginResp{
		Code:    200,
		Data:    resp.Token,
		Message: "登录成功",
	}, nil
}
