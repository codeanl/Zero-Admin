package user

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (*types.DeleteUserResp, error) {
	_, err := l.svcCtx.Sys.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除失败")
	}
	return &types.DeleteUserResp{
		Code:    200,
		Message: "删除用户成功",
	}, nil
}
