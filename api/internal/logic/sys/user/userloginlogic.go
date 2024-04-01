package user

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"google.golang.org/grpc/status"

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

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (*types.UserLoginResp, error) {
	resp, err := l.svcCtx.Sys.UserLogin(l.ctx, &sysclient.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.UserLoginResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
			return &types.UserLoginResp{
				Code:    400,
				Message: "系统出错",
			}, nil
		}
	}
	return &types.UserLoginResp{
		Code:    200,
		Data:    resp.Token,
		Message: "登录成功",
	}, nil
}
