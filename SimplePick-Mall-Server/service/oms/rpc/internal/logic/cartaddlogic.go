package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartAddLogic {
	return &CartAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加购物车
func (l *CartAddLogic) CartAdd(in *oms.CartAddReq) (*oms.CartAddResp, error) {
	role := &model.Cart{
		UserID: in.UserID,
		SkuID:  in.SkuID,
		Count:  in.Count,
	}
	_, err := l.svcCtx.CartModel.AddCart(role)
	if err != nil {
		return nil, err
	}
	return &oms.CartAddResp{}, nil
}
