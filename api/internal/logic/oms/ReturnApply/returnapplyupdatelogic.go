package ReturnApply

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyUpdateLogic {
	return &ReturnApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyUpdateLogic) ReturnApplyUpdate(req *types.UpdateReturnApplyReq) (resp *types.UpdateReturnApplyResp, err error) {
	_, err = l.svcCtx.Oms.ReturnApplyUpdate(l.ctx, &omsclient.ReturnApplyUpdateReq{
		ID:             req.ID,
		UserID:         req.UserID,
		OrderID:        req.OrderID,
		ReturnReasonID: req.ReturnReasonID,
		Status:         req.Status,
		Description:    req.Description,
		ProofPics:      req.ProofPics,
		ReturnAmount:   req.ReturnAmount,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateReturnApplyResp{
		Code:    200,
		Message: "更新用户成功",
	}, nil
}
