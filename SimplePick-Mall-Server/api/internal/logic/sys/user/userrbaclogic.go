package user

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRbacLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRbacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRbacLogic {
	return &UserRbacLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRbacLogic) UserRbac(req *types.UserRbacReq) (resp *types.UserRbacResp, err error) {
	_, err = l.svcCtx.Sys.UserRbac(l.ctx, &sysclient.UserRbacReq{
		ID:     req.ID,
		RoleID: req.RoleID,
	})
	if err != nil {
		return &types.UserRbacResp{
			Code:    400,
			Message: "分配失败",
		}, nil
	}
	return &types.UserRbacResp{
		Code:    200,
		Message: "分配成功",
	}, nil
}
