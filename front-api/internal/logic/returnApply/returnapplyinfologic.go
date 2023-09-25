package returnApply

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyInfoLogic {
	return &ReturnApplyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyInfoLogic) ReturnApplyInfo(req *types.ReturnApplyInfoReq) (*types.ReturnApplyInfoResp, error) {
	resp, _ := l.svcCtx.Oms.ReturnApplyInfo(l.ctx, &omsclient.ReturnApplyInfoReq{OrderID: req.OrderID})
	if resp.ID == 0 {
		return &types.ReturnApplyInfoResp{
			Code:    400,
			Message: "error",
		}, nil
	}
	data := types.ReturnApplyInfoData{
		ID:               resp.ID,
		UserID:           resp.UserID,
		OrderID:          resp.OrderID,
		ReturnReasonID:   resp.ReturnReasonID,
		ReturnReasonName: resp.ReturnReasonName,
		Status:           resp.Status,
		Description:      resp.Description,
		ProofPics:        resp.ProofPics,
		ReturnAmount:     resp.ReturnAmount,
	}
	return &types.ReturnApplyInfoResp{
		Code:    200,
		Message: "success",
		Data:    data,
	}, nil
}
