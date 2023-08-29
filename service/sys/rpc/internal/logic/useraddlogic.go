package logic

import (
	"SimplePick-Mall-Server/common/MD5"
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户
func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.UserAddResp, error) {
	//1.查询用户是否存在
	_, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if exist {
		return nil, errors.New("账户已存在")
	}
	//2.添加新用户
	admin := &model.User{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Password: MD5.SetPassword(in.Password),
		Gender:   in.Gender,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Status:   in.Status,
		CreateBy: in.CreateBy,
	}
	err := l.svcCtx.UserModel.AddUser(admin)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	//3.添加用户角色
	//for _, r := range in.RoleID {
	//	if r != 0 {
	//		err := l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
	//			UserID:   int64(admin.ID),
	//			RoleID:   r,
	//			CreateBy: in.CreateBy,
	//		})
	//		if err != nil {
	//		}
	//	}
	//}
	return &sys.UserAddResp{}, nil
}
