package member

import (
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdateLogic {
	return &MemberUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberUpdateLogic) MemberUpdate(req *types.UpdateMemberReq) (resp *types.UpdateMemberResp, err error) {
	_, err = l.svcCtx.Ums.MemberUpdate(l.ctx, &umsclient.MemberUpdateReq{
		Id:        req.Id,
		Username:  req.Username,
		Password:  req.Password,
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
		return &types.UpdateMemberResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateMemberResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
