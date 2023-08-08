package role

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {
	_, err = l.svcCtx.Sys.RoleAdd(l.ctx, &sysclient.RoleAddReq{
		Name:     req.Name,
		Remark:   req.Remark,
		CreateBy: l.ctx.Value("username").(string),
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
