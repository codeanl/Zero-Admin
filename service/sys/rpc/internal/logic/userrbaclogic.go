package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRbacLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRbacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRbacLogic {
	return &UserRbacLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分配权限
func (l *UserRbacLogic) UserRbac(in *sys.UserRbacReq) (*sys.UserRbacResp, error) {
	//把原有信息删除 添加新数据
	err := l.svcCtx.UserRoleModel.DeleteByUserID(in.ID)
	if err != nil {
		return nil, errors.New("修改用户角色失败")
	}
	//添加
	for _, r := range in.RoleID {
		if r != 0 {
			err := l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
				UserID: in.ID,
				RoleID: r,
			})
			if err != nil {
				return nil, errors.New("修改用户角色失败")
			}
		}
	}
	return &sys.UserRbacResp{}, nil
}
