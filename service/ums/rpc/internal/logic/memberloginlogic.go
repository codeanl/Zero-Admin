package logic

import (
	jwt "SimplePick-Mall-Server/common/JWT"
	"context"
	"errors"
	"time"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogic {
	return &MemberLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员登录
func (l *MemberLoginLogic) MemberLogin(in *ums.MemberLoginReq) (*ums.MemberLoginResp, error) {
	userInfo, exist, _ := l.svcCtx.MemberModel.GetMemberByMembername(in.Username)
	if !exist {
		return nil, errors.New("用户不存在")
	}
	if userInfo.Password != in.Password {
		return nil, errors.New("用户密码不正确")
	}
	info := &ums.MemberListData{
		Id:         int64(userInfo.ID),
		Username:   userInfo.Username,
		Password:   userInfo.Password,
		Nickname:   userInfo.Nickname,
		Phone:      userInfo.Phone,
		Status:     userInfo.Status,
		Avatar:     userInfo.Avatar,
		Gender:     userInfo.Gender,
		Email:      userInfo.Email,
		City:       userInfo.City,
		Job:        userInfo.Job,
		Signature:  userInfo.Signature,
		CreateTime: userInfo.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	AccessToken, err := jwt.GetMemberJwtToken(accessSecret, now, accessExpire, int64(userInfo.ID), userInfo.Username, userInfo.Nickname)
	if err != nil {
		return nil, err
	}
	resp := &ums.MemberLoginResp{
		Token:  AccessToken,
		Member: info,
	}
	return resp, nil
}
