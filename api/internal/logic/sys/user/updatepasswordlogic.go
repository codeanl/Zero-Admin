package user

import (
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
		return &types.UpdatePasswordResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdatePasswordResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
