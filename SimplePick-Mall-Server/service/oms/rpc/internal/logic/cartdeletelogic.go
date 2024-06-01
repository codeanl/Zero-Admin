package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartDeleteLogic {
	return &CartDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除购物车
func (l *CartDeleteLogic) CartDelete(in *oms.CartDeleteReq) (*oms.CartDeleteResp, error) {
	err := l.svcCtx.CartModel.DeleteCartByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &oms.CartDeleteResp{}, nil
}
