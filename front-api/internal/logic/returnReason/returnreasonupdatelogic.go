package returnReason

import (
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
