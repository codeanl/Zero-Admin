package logic

import (
	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"context"

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
func (l *UserInfoLogic) UserInfo(in *sys.InfoReq) (*sys.InfoResp, error) {
	//todo 设置假数据 后期修改

	//查询用户信息
	user, _ := l.svcCtx.UserModel.GetUserByID(in.ID)
	userInfo := &sys.UserInfo{
		ID:       int64(user.ID),
		Username: user.Username,
		Phone:    user.Phone,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   user.Status,
		RoleName: user.Username,
		CreateBy: user.CreateBy,
		UpdateBy: user.UpdateBy,
		CreateAt: user.CreateBy,
		UpdateAt: user.UpdateBy,
	}

	//获取用户的所有角色名字
	role, err := l.svcCtx.RoleModel.GetRoleByUserID(in.ID)
	if err != nil {
		return nil, err
	}
	var Roles []string
	for _, item := range role {
		Roles = append(Roles, item.Name)
	}
	//是否含有超级管理员这个角色
	targetRole := "超级管理员"
	contains := false
	for _, i := range Roles {
		if i == targetRole {
			contains = true
			break
		}
	}
	var Routes []string
	var Buttons []string
	//is
	if contains {
		menus, _, _ := l.svcCtx.MenuModel.GetMenuList()
		for _, i := range menus {
			if i.Type != "3" {
				Routes = append(Routes, i.TAG)
			} else {
				Buttons = append(Buttons, i.TAG)
			}
		}
	} else { //no
		menus, _ := l.svcCtx.MenuModel.GetMenusByUserID(in.ID)
		for _, i := range menus {
			if i.Type != "3" {
				Routes = append(Routes, i.TAG)
			} else {
				Buttons = append(Buttons, i.TAG)
			}
		}
	}

	return &sys.InfoResp{
		UserInfo: userInfo,
		Routes:   Routes,
		Roles:    Roles,
		Buttons:  Buttons,
	}, nil
}
