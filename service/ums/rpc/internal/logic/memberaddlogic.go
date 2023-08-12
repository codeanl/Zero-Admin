package logic

import (
	"SimplePick-Mall-Server/service/ums/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加会员
func (l *MemberAddLogic) MemberAdd(in *ums.MemberAddReq) (*ums.MemberAddResp, error) {
	_, exist, _ := l.svcCtx.MemberModel.GetMemberByMembername(in.Username)
	if exist {
		return nil, errors.New("账户已存在")
	}
	info := &model.Member{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Password: in.Password,
	}
	err := l.svcCtx.MemberModel.AddMember(info)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &ums.MemberAddResp{}, nil
}
