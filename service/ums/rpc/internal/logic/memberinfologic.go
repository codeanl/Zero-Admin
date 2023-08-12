package logic

import (
	"context"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberInfoLogic {
	return &MemberInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员详情
func (l *MemberInfoLogic) MemberInfo(in *ums.MemberInfoReq) (*ums.MemberInfoResp, error) {
	userInfo, err := l.svcCtx.MemberModel.GetMemberByID(in.Id)
	if err != nil {
		return nil, err
	}
	return &ums.MemberInfoResp{
		Id:       int64(userInfo.ID),
		Username: userInfo.Username,
		Nickname: userInfo.Nickname,
		Phone:    userInfo.Phone,
		Status:   userInfo.Status,
		Avatar:   userInfo.Avatar,
		Gender:   userInfo.Gender,
		Email:    userInfo.Email,
		City:     userInfo.City,
		Job:      userInfo.Job,
	}, nil
}
