package sku

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuUpdateLogic {
	return &SkuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuUpdateLogic) SkuUpdate(req *types.UpdateSkuReq) (resp *types.UpdateSkuResp, err error) {
	_, err = l.svcCtx.Pms.SkuUpdate(l.ctx, &pmsclient.SkuUpdateReq{
		ID:          req.ID,
		ProductID:   req.ProductID,
		Name:        req.Name,
		Pic:         req.Pic,
		Price:       req.Price,
		SkuSn:       req.SkuSn,
		Description: req.Description,
		Stock:       req.Stock,
		Tag:         req.TAG,
		SizeValueID: req.SizeValueID,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateSkuResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
