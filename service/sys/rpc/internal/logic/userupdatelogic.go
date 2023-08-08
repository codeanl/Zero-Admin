package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户
func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	err := l.svcCtx.UserModel.UpdateUser(in.ID, &model.User{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Gender:   in.Gender,
		Email:    in.Email,
		Status:   in.Status,
		UpdateBy: in.UpdateBy,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	//把原有信息删除 添加新数据
	//todo 后期获取原有数据 过滤添加删除
	//删除
	err = l.svcCtx.UserRoleModel.DeleteByUserID(in.ID)
	//添加
	for _, r := range in.RoleID {
		if r != 0 {
			err := l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
				UserID:   in.ID,
				RoleID:   r,
				CreateBy: in.UpdateBy,
			})
			if err != nil {
				// 处理错误，例如返回错误或进行日志记录
				return nil, errors.New("修改用户角色失败")
			}
		}
	}
	return &sys.UserUpdateResp{}, nil
}
