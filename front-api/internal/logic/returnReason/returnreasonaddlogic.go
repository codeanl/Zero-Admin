package returnReason

import (
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
