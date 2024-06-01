package logic

import (
	"SimplePick-Mall-Server/common/MD5"
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRestartPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartPasswordLogic {
	return &RestartPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置密码
func (l *RestartPasswordLogic) RestartPassword(in *sys.RestartPasswordReq) (*sys.RestartPasswordResp, error) {
	//默认密码为123456
	err := l.svcCtx.UserModel.UpdateUser(in.ID, &model.User{
		Password: MD5.SetPassword("123456"),
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &sys.RestartPasswordResp{}, nil
}
