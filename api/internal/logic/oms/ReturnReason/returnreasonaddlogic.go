package ReturnReason

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnReasonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonAddLogic {
	return &ReturnReasonAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnReasonAddLogic) ReturnReasonAdd(req *types.AddReturnReasonReq) (resp *types.AddReturnReasonResp, err error) {
	_, err = l.svcCtx.Oms.ReturnReasonAdd(l.ctx, &omsclient.ReturnReasonAddReq{
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		return &types.AddReturnReasonResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddReturnReasonResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
