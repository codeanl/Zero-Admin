package role

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.ListRoleReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.Sys.RoleList(l.ctx, &sysclient.RoleListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListRoleData
	for _, item := range resp.List {
		listUserData := types.ListRoleData{
			Id:       item.Id,
			Name:     item.Name,
			Remark:   item.Remark,
			CreateBy: item.CreateBy,
			CreateAt: item.CreateAt,
			UpdateBy: item.UpdateBy,
			UpdateAt: item.UpdateAt,
		}
		list = append(list, &listUserData)
	}

	return &types.ListRoleResp{
		Code:    200,
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
