package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

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
	routes := []string{"aaa", "User", "Category", "Discount", "ActivityEdit", "CouponRule",
		"Product", "Activity", "CouponAdd", "Trademark", "test1", "Attr", "ActivityAdd", "qqq", "CouponEdit",
		"OrderShow", "111", "Permission", "Spu", "UserList", "ClientUser", "Order", "33", "Coupon", "permision",
		"Acl", "ActivityRule", "Role", "RoleAuth", "222", "Refund", "1223", "OrderList", "Sku", "Log", "SysLog", "LoginLog"}
	//Roles := []string{"超级管理员"}
	Buttons := []string{"cuser.detail",
		"cuser.update",
		"cuser.delete",
		"btn.User.add",
		"btn.User.remove",
		"btn.User.update",
		"btn.User.assgin",
		"btn.Role.assgin",
		"btn.Role.add",
		"btn.Role.update",
		"btn.Role.remove",
		"btn.Permission.add",
		"btn.Permission.update",
		"btn.Permission.remove",
		"btn.Activity.add",
		"btn.Activity.update",
		"btn.Activity.rule",
		"btn.Coupon.add",
		"btn.Coupon.update",
		"btn.Coupon.rule",
		"btn.OrderList.detail",
		"btn.OrderList.Refund",
		"btn.UserList.lock",
		"btn.Category.add",
		"btn.Category.update",
		"btn.Category.remove",
		"btn.Trademark.add",
		"btn.Trademark.update",
		"btn.Trademark.remove",
		"btn.Attr.add",
		"btn.Attr.update",
		"btn.Attr.remove",
		"btn.Spu.add",
		"btn.Spu.addsku",
		"btn.Spu.update",
		"btn.Spu.skus",
		"btn.Spu.delete",
		"btn.Sku.updown",
		"btn.Sku.update",
		"btn.Sku.detail",
		"btn.Sku.remove",
		"btn.all",
		"tuiguang",
		"btn.test.2",
		"fdfg"}
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
	//获取用户的所有角色
	role, err := l.svcCtx.RoleModel.GetRoleByUserID(in.ID)
	if err != nil {
		return nil, err
	}
	var Roles []string
	for _, item := range role {
		Roles = append(Roles, item.Name)
	}
	return &sys.InfoResp{
		UserInfo: userInfo,
		Routes:   routes,
		Roles:    Roles,
		Buttons:  Buttons,
	}, nil
}
