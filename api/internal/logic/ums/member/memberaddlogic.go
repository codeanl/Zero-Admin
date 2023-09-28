package member

import (
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req *types.AddMemberReq) (resp *types.AddMemberResp, err error) {
	_, err = l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Phone:     req.Phone,
		Status:    req.Status,
		Avatar:    req.Avatar,
		Gender:    req.Gender,
		Email:     req.Email,
		City:      req.City,
		Job:       req.Job,
		Signature: req.Signature,
	})
	if err != nil {
		return &types.AddMemberResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddMemberResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
