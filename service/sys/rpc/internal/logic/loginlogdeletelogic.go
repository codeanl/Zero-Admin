package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogDeleteLogic {
	return &LoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除登录日志
func (l *LoginLogDeleteLogic) LoginLogDelete(in *sys.LoginLogDeleteReq) (*sys.LoginLogDeleteResp, error) {
	err := l.svcCtx.LoginLogModel.DeleteLoginLogByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sys.LoginLogDeleteResp{}, nil
}
