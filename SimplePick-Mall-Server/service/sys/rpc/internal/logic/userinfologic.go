package logic

import (
	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户信息
func (l *UserInfoLogic) UserInfo(in *sys.UserInfoReq) (*sys.UserInfoResp, error) {
	//查询用户信息
	user, _ := l.svcCtx.UserModel.GetUserByID(in.Id)

	var userInfo *sys.UserInfo
	jsonData, err := json.Marshal(user)
	err = json.Unmarshal(jsonData, &userInfo)

	//获取用户的所有角色名字
	role, err := l.svcCtx.RoleModel.GetRoleByUserID(in.Id)
	if err != nil {
		return nil, err
	}
	var Roles []string
	for _, item := range role {
		Roles = append(Roles, item.Name)
	}
	//是否含有超级管理员这个角色 超级管理员原有所有的菜单权限
	targetRole := "超级管理员"
	contains := false
	for _, i := range Roles {
		if i == targetRole {
			contains = true
			break
		}
	}
	var Routes []string
	//超级管理员 获取所有得到菜单
	if contains {
		menus, _, _ := l.svcCtx.MenuModel.GetMenuList()
		for _, i := range menus {
			Routes = append(Routes, i.TAG)
		}
	} else {
		//其他角色  通过用户id-->查询角色-->查询角色原拥有菜单
		menus, _ := l.svcCtx.MenuModel.GetMenusByUserID(in.Id)
		for _, i := range menus {
			Routes = append(Routes, i.TAG)
		}
	}
	return &sys.UserInfoResp{
		UserInfo: userInfo,
		Routes:   Routes,
		Roles:    Roles,
	}, nil
}
