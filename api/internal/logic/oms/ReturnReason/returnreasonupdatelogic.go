package ReturnReason

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnReasonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonUpdateLogic {
	return &ReturnReasonUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnReasonUpdateLogic) ReturnReasonUpdate(req *types.UpdateReturnReasonReq) (resp *types.UpdateReturnReasonResp, err error) {
	_, err = l.svcCtx.Oms.ReturnReasonUpdate(l.ctx, &omsclient.ReturnReasonUpdateReq{
		ID:     req.ID,
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		return &types.UpdateReturnReasonResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateReturnReasonResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
