package user

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	_, err = l.svcCtx.Sys.UserUpdate(l.ctx, &sysclient.UserUpdateReq{
		ID:       req.ID,
		Username: req.Username,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Email:    req.Email,
		Status:   req.Status,
		UpdateBy: l.ctx.Value("username").(string),
		RoleID:   req.RoleID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateUserResp{
		Code:    200,
		Message: "成功",
	}, nil
}
