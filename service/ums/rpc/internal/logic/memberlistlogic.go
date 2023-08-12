package logic

import (
	"context"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员登录
func (l *MemberListLogic) MemberList(in *ums.MemberListReq) (*ums.MemberListResp, error) {
	all, total, err := l.svcCtx.MemberModel.GetMemberList(in)
	if err != nil {
		return nil, err
	}
	var list []*ums.MemberListData
	for _, user := range all {
		list = append(list, &ums.MemberListData{
			Id:         int64(user.ID),
			Username:   user.Username,
			Password:   user.Password,
			Nickname:   user.Nickname,
			Phone:      user.Phone,
			Status:     user.Status,
			Avatar:     user.Avatar,
			Gender:     user.Gender,
			Email:      user.Email,
			City:       user.City,
			Job:        user.Job,
			Signature:  user.Signature,
			CreateTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &ums.MemberListResp{
		Total: total,
		List:  list,
	}, nil
}
