package role

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleMenuLogic {
	return &UpdateRoleMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleMenuLogic) UpdateRoleMenu(req *types.UpdateRoleMenuReq) (*types.UpdateRoleMenuResp, error) {
	_, err := l.svcCtx.Sys.UpdateMenuRole(l.ctx, &sysclient.UpdateMenuRoleReq{
		RoleId:   req.RoleId,
		MenuIds:  req.MenuIds,
		CreateBy: l.ctx.Value("username").(string),
	})
	if err != nil {
		return &types.UpdateRoleMenuResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateRoleMenuResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
