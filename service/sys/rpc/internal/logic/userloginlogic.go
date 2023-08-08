package logic

import (
	jwt "SimplePick-Mall-Server/common/JWT"
	"SimplePick-Mall-Server/common/MD5"
	"context"
	"errors"
	"time"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *UserLoginLogic) UserLogin(in *sys.LoginReq) (*sys.LoginResp, error) {
	user, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if !exist {
		return nil, errors.New("用户不存在")
	}
	//校验密码
	yes := MD5.CheckPassword(user.Password, in.Password)
	if !yes {
		return nil, errors.New("用户密码不正确")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	AccessToken, err := jwt.GetJwtToken(accessSecret, now, accessExpire, int64(user.ID), user.Username, user.Nickname)
	if err != nil {
		return nil, err
	}
	return &sys.LoginResp{
		Token:  AccessToken,
		UserID: int64(user.ID),
	}, nil
}
