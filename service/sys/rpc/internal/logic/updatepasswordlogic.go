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

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新密码
func (l *UpdatePasswordLogic) UpdatePassword(in *sys.UpdatePasswordReq) (*sys.UpdatePasswordResp, error) {
	//获取信息得到密码
	my, _ := l.svcCtx.UserModel.GetUserByID(in.ID)
	yes := MD5.CheckPassword(my.Password, in.OldPassword)
	if !yes {
		return nil, errors.New("旧密码不正确")
	}
	//更新
	err := l.svcCtx.UserModel.UpdateUser(in.ID, &model.User{
		Password: MD5.SetPassword(in.NewPassword),
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &sys.UpdatePasswordResp{}, nil
}
