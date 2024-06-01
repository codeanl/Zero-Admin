package ReturnReason

import (
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonListLogic {
	return &ReturnReasonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnReasonListLogic) ReturnReasonList(req *types.ListReturnReasonReq) (resp *types.ListReturnReasonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
