package user

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.UpdatePasswordResp, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err = l.svcCtx.Sys.UpdatePassword(l.ctx, &sysclient.UpdatePasswordReq{
		ID:          id,
		NewPassword: req.NewPassword,
		OldPassword: req.OldPassword,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新密码失败")
	}
	return &types.UpdatePasswordResp{
		Code:    200,
		Message: "成功",
	}, nil
	return
}
