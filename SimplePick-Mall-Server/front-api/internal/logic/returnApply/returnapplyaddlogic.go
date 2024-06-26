package returnApply

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyAddLogic {
	return &ReturnApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyAddLogic) ReturnApplyAdd(req *types.AddReturnApplyReq) (resp *types.AddReturnApplyResp, err error) {
	order, _ := l.svcCtx.Oms.OrderInfo(l.ctx, &omsclient.OrderInfoReq{Id: req.OrderID})
	_, err = l.svcCtx.Oms.ReturnApplyAdd(l.ctx, &omsclient.ReturnApplyAddReq{
		UserID:         req.UserID,
		OrderID:        req.OrderID,
		ReturnReasonID: req.ReturnReasonID,
		Status:         req.Status,
		Description:    req.Description,
		ProofPics:      req.ProofPics,
		ReturnAmount:   req.ReturnAmount,
		MerchantID:     order.OrderInfo.MerchantID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddReturnApplyResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
