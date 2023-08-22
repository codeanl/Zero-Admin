package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleLogic {
	return &UpdateMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色拥有菜单
func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *sys.UpdateMenuRoleReq) (*sys.UpdateMenuRoleResp, error) {
	//删除当前角色拥有的菜单 重新添加
	_ = l.svcCtx.RoleMenuModel.DeleteByRoleId(in.RoleId)
	for _, menuId := range in.MenuIds {
		//重新添加
		_ = l.svcCtx.RoleMenuModel.AddRoleMenu(&model.RoleMenu{
			RoleID:   in.RoleId,
			MenuID:   menuId,
			CreateBy: in.CreateBy,
		})
	}
	return &sys.UpdateMenuRoleResp{}, nil
}
