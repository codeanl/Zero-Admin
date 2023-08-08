package role

import (
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

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

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req *types.RoleMenuReq) (resp *types.RoleMenuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
