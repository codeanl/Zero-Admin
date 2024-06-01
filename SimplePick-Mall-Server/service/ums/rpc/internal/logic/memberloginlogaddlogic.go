package logic

import (
	"SimplePick-Mall-Server/service/ums/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogAddLogic {
	return &MemberLoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加登录日志
func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(in *ums.MemberLoginLogAddReq) (*ums.MemberLoginLogAddResp, error) {
	info := &model.MemberLoginLog{
		UserID: in.UserID,
		IP:     in.IP,
	}
	err := l.svcCtx.MemberLoginLogModel.AddMemberLoginLog(info)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &ums.MemberLoginLogAddResp{}, nil
}
