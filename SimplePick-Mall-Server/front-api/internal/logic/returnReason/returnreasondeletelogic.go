package returnReason

import (
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnReasonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonDeleteLogic {
	return &ReturnReasonDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnReasonDeleteLogic) ReturnReasonDelete(req *types.DeleteReturnReasonReq) (resp *types.DeleteReturnReasonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
