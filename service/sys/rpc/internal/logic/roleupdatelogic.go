package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色
func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	err := l.svcCtx.RoleModel.UpdateRole(in.Id, &model.Role{
		Name:     in.Name,
		Remark:   in.Remark,
		UpdateBy: in.UpdateBy,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &sys.RoleUpdateResp{}, nil
}
