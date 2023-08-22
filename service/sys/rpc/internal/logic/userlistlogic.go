package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户列表
func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {
	//查询所有的用户
	all, total, err := l.svcCtx.UserModel.GetUserList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.UserList
	for _, user := range all {
		//获取用户的所有角色
		role, err := l.svcCtx.RoleModel.GetRoleByUserID(int64(user.ID))
		if err != nil {
			return nil, err
		}
		var Roles []string
		for _, item := range role {
			Roles = append(Roles, item.Name)
		}
		list = append(list, &sys.UserList{
			Id:       int64(user.ID),
			Username: user.Username,
			Phone:    user.Phone,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Status:   user.Status,
			CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			CreateBy: user.CreateBy,
			UpdateBy: user.UpdateBy,
			RoleName: Roles,
		})
	}
	return &sys.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
