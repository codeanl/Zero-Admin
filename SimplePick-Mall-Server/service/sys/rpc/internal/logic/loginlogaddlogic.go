package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加登录日志
func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	err := l.svcCtx.LoginLogModel.AddLoginLog(&model.LoginLog{
		UserID: in.UserID,
		IP:     in.IP,
	})
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &sys.LoginLogAddResp{}, nil
}
