package auth

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		return nil, errorx.NewDefaultError("更新失败")
	}
	return &types.UpdateMemberResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
