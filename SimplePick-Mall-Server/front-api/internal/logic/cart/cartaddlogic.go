package cart

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartAddLogic {
	return &CartAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartAddLogic) CartAdd(req *types.AddCartReq) (resp *types.AddCartResp, err error) {
	_, err = l.svcCtx.Oms.CartAdd(l.ctx, &omsclient.CartAddReq{
		UserID: req.UserID,
		SkuID:  req.SkuID,
		Count:  req.Count,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddCartResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
