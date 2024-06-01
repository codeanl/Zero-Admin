package logic

import (
	"SimplePick-Mall-Server/service/ums/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdateLogic {
	return &MemberUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员更新
func (l *MemberUpdateLogic) MemberUpdate(in *ums.MemberUpdateReq) (*ums.MemberUpdateResp, error) {
	err := l.svcCtx.MemberModel.UpdateMember(in.Id, &model.Member{
		Username:  in.Username,
		Password:  in.Password,
		Nickname:  in.Nickname,
		Phone:     in.Phone,
		Status:    in.Status,
		Avatar:    in.Avatar,
		Gender:    in.Gender,
		Email:     in.Email,
		City:      in.City,
		Job:       in.Job,
		Signature: in.Signature,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &ums.MemberUpdateResp{}, nil
}
