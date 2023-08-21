package ReturnReason

import (
	"SimplePick-Mall-Server/common/errorx"
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
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeleteReturnReasonResp{
		Code:    200,
		Message: "删除用户成功",
	}, nil
}
