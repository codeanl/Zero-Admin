package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新购物车
func (l *CartUpdateLogic) CartUpdate(in *oms.CartUpdateReq) (*oms.CartUpdateResp, error) {
	err := l.svcCtx.CartModel.UpdateCart(in.ID, &model.Cart{
		UserID: in.UserID,
		SkuID:  in.SkuID,
		Count:  in.Count,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &oms.CartUpdateResp{}, nil
}
