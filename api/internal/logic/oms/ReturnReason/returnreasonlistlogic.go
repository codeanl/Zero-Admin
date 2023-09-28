package ReturnReason

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

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

func (l *ReturnReasonListLogic) ReturnReasonList(req *types.ListReturnReasonReq) (*types.ListReturnReasonResp, error) {
	resp, err := l.svcCtx.Oms.ReturnReasonList(l.ctx, &omsclient.ReturnReasonListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListReturnReasonResp{
			Code:    400,
			Message: "查询失败",
		}, nil

	}
	var list []types.ListReturnReasonData
	for _, item := range resp.List {
		listUserData := types.ListReturnReasonData{
			ID:     item.ID,
			Name:   item.Name,
			Status: item.Status,
		}
		list = append(list, listUserData)
	}
	return &types.ListReturnReasonResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
