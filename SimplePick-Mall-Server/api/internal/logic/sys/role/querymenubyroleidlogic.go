package role

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req *types.RoleMenuReq) (*types.RoleMenuResp, error) {
	resp, err := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{})
	if err != nil {
		return &types.RoleMenuResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var listIds []int64
	for _, i := range resp.MenuList {
		listIds = append(listIds, i.ID)
	}
	//如果角色不是admin则根据roleId查询菜单
	if req.Id != 1 {
		ids, _ := l.svcCtx.Sys.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{Id: req.Id})
		listIds = ids.Ids
	}
	return &types.RoleMenuResp{
		Data:    listIds,
		Code:    200,
		Message: "查询成功",
	}, nil
}
