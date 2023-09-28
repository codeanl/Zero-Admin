package user

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (*types.AddUserResp, error) {
	resp, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Username: req.Username,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Password: "123456",
		Gender:   req.Gender,
		Avatar:   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif", //默认头像
		Email:    req.Email,
		Status:   req.Status,
		CreateBy: l.ctx.Value("username").(string),
		RoleID:   req.RoleID,
	})
	if err != nil {
		return &types.AddUserResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddUserResp{
		Code:    200,
		Message: "添加成功",
		Data:    resp.ID,
	}, nil
}
