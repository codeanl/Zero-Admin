package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 角色列表
func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	all, total, err := l.svcCtx.RoleModel.GetRoleList(in)
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
	return &sys.RoleListResp{
		Total: total,
		List:  list,
	}, nil
}
