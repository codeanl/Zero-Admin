// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/logic"
	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

// 用户登录
func (s *SysServer) UserLogin(ctx context.Context, in *sys.LoginReq) (*sys.LoginResp, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

// 用户信息
func (s *SysServer) UserInfo(ctx context.Context, in *sys.InfoReq) (*sys.InfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

// 添加用户
func (s *SysServer) UserAdd(ctx context.Context, in *sys.UserAddReq) (*sys.UserAddResp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

// 更新用户
func (s *SysServer) UserUpdate(ctx context.Context, in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

// 删除用户
func (s *SysServer) UserDelete(ctx context.Context, in *sys.UserDeleteReq) (*sys.UserDeleteResp, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

// 用户列表
func (s *SysServer) UserList(ctx context.Context, in *sys.UserListReq) (*sys.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

// 更新密码
func (s *SysServer) UpdatePassword(ctx context.Context, in *sys.UpdatePasswordReq) (*sys.UpdatePasswordResp, error) {
	l := logic.NewUpdatePasswordLogic(ctx, s.svcCtx)
	return l.UpdatePassword(in)
}

// 重置密码
func (s *SysServer) RestartPassword(ctx context.Context, in *sys.RestartPasswordReq) (*sys.RestartPasswordResp, error) {
	l := logic.NewRestartPasswordLogic(ctx, s.svcCtx)
	return l.RestartPassword(in)
}

// 添加登录日志
func (s *SysServer) LoginLogAdd(ctx context.Context, in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	l := logic.NewLoginLogAddLogic(ctx, s.svcCtx)
	return l.LoginLogAdd(in)
}

// 登录日志列表
func (s *SysServer) LoginLogList(ctx context.Context, in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

// 删除登录日志
func (s *SysServer) LoginLogDelete(ctx context.Context, in *sys.LoginLogDeleteReq) (*sys.LoginLogDeleteResp, error) {
	l := logic.NewLoginLogDeleteLogic(ctx, s.svcCtx)
	return l.LoginLogDelete(in)
}

// 添加自提点
func (s *SysServer) PlaceAdd(ctx context.Context, in *sys.PlaceAddReq) (*sys.PlaceAddResp, error) {
	l := logic.NewPlaceAddLogic(ctx, s.svcCtx)
	return l.PlaceAdd(in)
}

// 自提点列表
func (s *SysServer) PlaceList(ctx context.Context, in *sys.PlaceListReq) (*sys.PlaceListResp, error) {
	l := logic.NewPlaceListLogic(ctx, s.svcCtx)
	return l.PlaceList(in)
}

// 更新自提点
func (s *SysServer) PlaceUpdate(ctx context.Context, in *sys.PlaceUpdateReq) (*sys.PlaceUpdateResp, error) {
	l := logic.NewPlaceUpdateLogic(ctx, s.svcCtx)
	return l.PlaceUpdate(in)
}

// 删除自提点
func (s *SysServer) PlaceDelete(ctx context.Context, in *sys.PlaceDeleteReq) (*sys.PlaceDeleteResp, error) {
	l := logic.NewPlaceDeleteLogic(ctx, s.svcCtx)
	return l.PlaceDelete(in)
}

// 添加角色
func (s *SysServer) RoleAdd(ctx context.Context, in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	l := logic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

// 更新角色
func (s *SysServer) RoleUpdate(ctx context.Context, in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	l := logic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

// 删除角色
func (s *SysServer) RoleDelete(ctx context.Context, in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	l := logic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

// 角色列表
func (s *SysServer) RoleList(ctx context.Context, in *sys.RoleListReq) (*sys.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

// 用户拥有的角色列表
func (s *SysServer) RoleByUserList(ctx context.Context, in *sys.RoleByUserListReq) (*sys.RoleByUserListResp, error) {
	l := logic.NewRoleByUserListLogic(ctx, s.svcCtx)
	return l.RoleByUserList(in)
}

// 获取角色的菜单
func (s *SysServer) QueryMenuByRoleId(ctx context.Context, in *sys.QueryMenuByRoleIdReq) (*sys.QueryMenuByRoleIdResp, error) {
	l := logic.NewQueryMenuByRoleIdLogic(ctx, s.svcCtx)
	return l.QueryMenuByRoleId(in)
}

// 更新角色拥有菜单
func (s *SysServer) UpdateMenuRole(ctx context.Context, in *sys.UpdateMenuRoleReq) (*sys.UpdateMenuRoleResp, error) {
	l := logic.NewUpdateMenuRoleLogic(ctx, s.svcCtx)
	return l.UpdateMenuRole(in)
}

// 添加菜单
func (s *SysServer) MenuAdd(ctx context.Context, in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	l := logic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

// 菜单列表
func (s *SysServer) MenuList(ctx context.Context, in *sys.MenuListReq) (*sys.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

// 更新菜单
func (s *SysServer) MenuUpdate(ctx context.Context, in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

// 删除菜单
func (s *SysServer) MenuDelete(ctx context.Context, in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

// 添加日志
func (s *SysServer) SysLogAdd(ctx context.Context, in *sys.LogAddReq) (*sys.LogAddResp, error) {
	l := logic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

// 日志列表
func (s *SysServer) SysLogList(ctx context.Context, in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	l := logic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

// 删除日志
func (s *SysServer) SysLogDelete(ctx context.Context, in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	l := logic.NewSysLogDeleteLogic(ctx, s.svcCtx)
	return l.SysLogDelete(in)
}