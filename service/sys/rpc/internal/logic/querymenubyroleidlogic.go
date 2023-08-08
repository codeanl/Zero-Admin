package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取角色的菜单
func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *sys.QueryMenuByRoleIdReq) (*sys.QueryMenuByRoleIdResp, error) {

	return &sys.QueryMenuByRoleIdResp{
		Ids: nil,
	}, nil
}
