package ReturnReason

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

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
	_, err = l.svcCtx.Oms.ReturnReasonDelete(l.ctx, &omsclient.ReturnReasonDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteReturnReasonResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteReturnReasonResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
