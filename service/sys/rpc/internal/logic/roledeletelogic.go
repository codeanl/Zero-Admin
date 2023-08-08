package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *RoleDeleteLogic) RoleDelete(in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	err := l.svcCtx.RoleModel.DeleteRoleByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sys.RoleDeleteResp{}, nil
}
