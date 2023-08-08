package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleByUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleByUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleByUserListLogic {
	return &RoleByUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户拥有的角色列表
func (l *RoleByUserListLogic) RoleByUserList(in *sys.RoleByUserListReq) (*sys.RoleByUserListResp, error) {
	all, err := l.svcCtx.RoleModel.GetRoleByUserList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.RoleListData
	for _, role := range all {
		list = append(list, &sys.RoleListData{
			Id:       int64(role.ID),
			Name:     role.Name,
			Remark:   role.Remark,
			CreateBy: role.CreateBy,
			CreateAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateBy: role.UpdateBy,
			UpdateAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &sys.RoleByUserListResp{
		List: list,
	}, nil
}
