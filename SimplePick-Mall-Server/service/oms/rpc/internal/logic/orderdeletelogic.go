package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新订单
func (l *OrderDeleteLogic) OrderDelete(in *oms.OrderDeleteReq) (*oms.OrderDeleteResp, error) {
	err := l.svcCtx.OrderModel.DeleteOrderByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &oms.OrderDeleteResp{}, nil
}
