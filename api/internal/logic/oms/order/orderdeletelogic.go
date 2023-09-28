package order

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req *types.DeleteOrderReq) (resp *types.DeleteOrderResp, err error) {
	_, err = l.svcCtx.Oms.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteOrderResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteOrderResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
