package logic

import (
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 购物车列表
func (l *CartListLogic) CartList(in *oms.CartListReq) (*oms.CartListResp, error) {
	all, total, err := l.svcCtx.CartModel.GetCartList(in)
	if err != nil {
		return nil, err
	}
	var list []*oms.CartListData
	for _, role := range all {
		list = append(list, &oms.CartListData{
			ID:     int64(role.ID),
			UserID: role.UserID,
			SkuID:  role.SkuID,
			Count:  role.Count,
		})
	}
	return &oms.CartListResp{
		Total: total,
		List:  list,
	}, nil
}
