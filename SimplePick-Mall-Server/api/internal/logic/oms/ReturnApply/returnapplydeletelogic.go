package ReturnApply

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyDeleteLogic {
	return &ReturnApplyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyDeleteLogic) ReturnApplyDelete(req *types.DeleteReturnApplyReq) (resp *types.DeleteReturnApplyResp, err error) {
	_, err = l.svcCtx.Oms.ReturnApplyDelete(l.ctx, &omsclient.ReturnApplyDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteReturnApplyResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteReturnApplyResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
