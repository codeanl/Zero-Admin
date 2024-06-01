package logic

import (
	jwt "SimplePick-Mall-Server/common/JWT"
	"SimplePick-Mall-Server/common/MD5"
	"SimplePick-Mall-Server/common/ip"
	"SimplePick-Mall-Server/service/sys/model"
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

//UserLogin 用户登录
func (l *UserLoginLogic) UserLogin(in *sys.UserLoginReq) (*sys.UserLoginResp, error) {
	//查询用户是否存在
	user, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if !exist {
		return nil, errors.New("用户不存在")
	}
	//查询用户是否锁定
	if user.Status == "0" {
		return nil, errors.New("用户已锁定")
	}
	//校验密码
	yes := MD5.CheckPassword(user.Password, in.Password)
	if !yes {
		return nil, errors.New("用户密码不正确")
	}
	//生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	Token, err := jwt.GetJwtToken(accessSecret, now, accessExpire, int64(user.ID), user.Username, user.Nickname)
	if err != nil {
		return nil, err
	}
	//添加ip登录日志
	_ = l.svcCtx.LoginLogModel.AddLoginLog(&model.LoginLog{
		UserID: int64(user.ID),
		IP:     ip.GetOutboundIP(),
	})
	//返回响应
	return &sys.UserLoginResp{
		Token:  Token,
		UserId: int64(user.ID),
	}, nil
}
