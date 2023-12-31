// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package sysclient

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq               = sys.InfoReq
	InfoResp              = sys.InfoResp
	LogAddReq             = sys.LogAddReq
	LogAddResp            = sys.LogAddResp
	LoginLogAddReq        = sys.LoginLogAddReq
	LoginLogAddResp       = sys.LoginLogAddResp
	LoginLogDeleteReq     = sys.LoginLogDeleteReq
	LoginLogDeleteResp    = sys.LoginLogDeleteResp
	LoginLogListData      = sys.LoginLogListData
	LoginLogListReq       = sys.LoginLogListReq
	LoginLogListResp      = sys.LoginLogListResp
	LoginReq              = sys.LoginReq
	LoginResp             = sys.LoginResp
	MenuAddReq            = sys.MenuAddReq
	MenuAddResp           = sys.MenuAddResp
	MenuDeleteReq         = sys.MenuDeleteReq
	MenuDeleteResp        = sys.MenuDeleteResp
	MenuList              = sys.MenuList
	MenuListReq           = sys.MenuListReq
	MenuListResp          = sys.MenuListResp
	MenuUpdateReq         = sys.MenuUpdateReq
	MenuUpdateResp        = sys.MenuUpdateResp
	PlaceAddReq           = sys.PlaceAddReq
	PlaceAddResp          = sys.PlaceAddResp
	PlaceDeleteReq        = sys.PlaceDeleteReq
	PlaceDeleteResp       = sys.PlaceDeleteResp
	PlaceInfoReq          = sys.PlaceInfoReq
	PlaceInfoResp         = sys.PlaceInfoResp
	PlaceListData         = sys.PlaceListData
	PlaceListReq          = sys.PlaceListReq
	PlaceListResp         = sys.PlaceListResp
	PlaceUpdateReq        = sys.PlaceUpdateReq
	PlaceUpdateResp       = sys.PlaceUpdateResp
	QueryMenuByRoleIdReq  = sys.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = sys.QueryMenuByRoleIdResp
	RestartPasswordReq    = sys.RestartPasswordReq
	RestartPasswordResp   = sys.RestartPasswordResp
	RoleAddReq            = sys.RoleAddReq
	RoleAddResp           = sys.RoleAddResp
	RoleByUserListReq     = sys.RoleByUserListReq
	RoleByUserListResp    = sys.RoleByUserListResp
	RoleDeleteReq         = sys.RoleDeleteReq
	RoleDeleteResp        = sys.RoleDeleteResp
	RoleListData          = sys.RoleListData
	RoleListReq           = sys.RoleListReq
	RoleListResp          = sys.RoleListResp
	RoleUpdateReq         = sys.RoleUpdateReq
	RoleUpdateResp        = sys.RoleUpdateResp
	SysLogDeleteReq       = sys.SysLogDeleteReq
	SysLogDeleteResp      = sys.SysLogDeleteResp
	SysLogListData        = sys.SysLogListData
	SysLogListReq         = sys.SysLogListReq
	SysLogListResp        = sys.SysLogListResp
	UpdateMenuRoleReq     = sys.UpdateMenuRoleReq
	UpdateMenuRoleResp    = sys.UpdateMenuRoleResp
	UpdatePasswordReq     = sys.UpdatePasswordReq
	UpdatePasswordResp    = sys.UpdatePasswordResp
	UserAddReq            = sys.UserAddReq
	UserAddResp           = sys.UserAddResp
	UserDeleteReq         = sys.UserDeleteReq
	UserDeleteResp        = sys.UserDeleteResp
	UserInfo              = sys.UserInfo
	UserList              = sys.UserList
	UserListReq           = sys.UserListReq
	UserListResp          = sys.UserListResp
	UserUpdateReq         = sys.UserUpdateReq
	UserUpdateResp        = sys.UserUpdateResp

	Sys interface {
		// 用户登录
		UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 用户信息
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		// 添加用户
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		// 更新用户
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		// 删除用户
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		// 用户列表
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		// 更新密码
		UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
		// 重置密码
		RestartPassword(ctx context.Context, in *RestartPasswordReq, opts ...grpc.CallOption) (*RestartPasswordResp, error)
		// 添加登录日志
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
		// 登录日志列表
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		// 删除登录日志
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		// 添加自提点
		PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error)
		// 自提点列表
		PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error)
		// 更新自提点
		PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error)
		// 删除自提点
		PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error)
		// 自提点详情
		PlaceInfo(ctx context.Context, in *PlaceInfoReq, opts ...grpc.CallOption) (*PlaceInfoResp, error)
		// 添加角色
		RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error)
		// 更新角色
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error)
		// 删除角色
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		// 角色列表
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		// 用户拥有的角色列表
		RoleByUserList(ctx context.Context, in *RoleByUserListReq, opts ...grpc.CallOption) (*RoleByUserListResp, error)
		// 获取角色的菜单
		QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error)
		// 更新角色拥有菜单
		UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error)
		// 添加菜单
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		// 菜单列表
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		// 更新菜单
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		// 删除菜单
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
		// 添加日志
		SysLogAdd(ctx context.Context, in *LogAddReq, opts ...grpc.CallOption) (*LogAddResp, error)
		// 日志列表
		SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error)
		// 删除日志
		SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

// 用户登录
func (m *defaultSys) UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// 用户信息
func (m *defaultSys) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

// 添加用户
func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

// 更新用户
func (m *defaultSys) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

// 删除用户
func (m *defaultSys) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

// 用户列表
func (m *defaultSys) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

// 更新密码
func (m *defaultSys) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdatePassword(ctx, in, opts...)
}

// 重置密码
func (m *defaultSys) RestartPassword(ctx context.Context, in *RestartPasswordReq, opts ...grpc.CallOption) (*RestartPasswordResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RestartPassword(ctx, in, opts...)
}

// 添加登录日志
func (m *defaultSys) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in, opts...)
}

// 登录日志列表
func (m *defaultSys) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

// 删除登录日志
func (m *defaultSys) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

// 添加自提点
func (m *defaultSys) PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PlaceAdd(ctx, in, opts...)
}

// 自提点列表
func (m *defaultSys) PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PlaceList(ctx, in, opts...)
}

// 更新自提点
func (m *defaultSys) PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PlaceUpdate(ctx, in, opts...)
}

// 删除自提点
func (m *defaultSys) PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PlaceDelete(ctx, in, opts...)
}

// 自提点详情
func (m *defaultSys) PlaceInfo(ctx context.Context, in *PlaceInfoReq, opts ...grpc.CallOption) (*PlaceInfoResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.PlaceInfo(ctx, in, opts...)
}

// 添加角色
func (m *defaultSys) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

// 更新角色
func (m *defaultSys) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

// 删除角色
func (m *defaultSys) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

// 角色列表
func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

// 用户拥有的角色列表
func (m *defaultSys) RoleByUserList(ctx context.Context, in *RoleByUserListReq, opts ...grpc.CallOption) (*RoleByUserListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleByUserList(ctx, in, opts...)
}

// 获取角色的菜单
func (m *defaultSys) QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.QueryMenuByRoleId(ctx, in, opts...)
}

// 更新角色拥有菜单
func (m *defaultSys) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UpdateMenuRole(ctx, in, opts...)
}

// 添加菜单
func (m *defaultSys) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

// 菜单列表
func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

// 更新菜单
func (m *defaultSys) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

// 删除菜单
func (m *defaultSys) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

// 添加日志
func (m *defaultSys) SysLogAdd(ctx context.Context, in *LogAddReq, opts ...grpc.CallOption) (*LogAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogAdd(ctx, in, opts...)
}

// 日志列表
func (m *defaultSys) SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogList(ctx, in, opts...)
}

// 删除日志
func (m *defaultSys) SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SysLogDelete(ctx, in, opts...)
}
