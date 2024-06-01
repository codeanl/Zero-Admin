package role

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleByUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleByUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleByUserListLogic {
	return &RoleByUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleByUserListLogic) RoleByUserList(req *types.RoleByUserListReq) (*types.RoleByUserListResp, error) {
	resp, err := l.svcCtx.Sys.RoleByUserList(l.ctx, &sysclient.RoleByUserListReq{
		UserID: req.UserID,
	})
	if err != nil {
		return &types.RoleByUserListResp{
			Code:    400,
			Message: "查询失败",
		}, nil
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

	return &types.RoleByUserListResp{
		Code:    200,
		Message: "查询成功",
		Data:    list,
	}, nil
}
