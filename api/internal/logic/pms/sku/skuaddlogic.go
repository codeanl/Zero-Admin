package sku

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuAddLogic {
	return &SkuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuAddLogic) SkuAdd(req *types.AddSkuReq) (resp *types.AddSkuResp, err error) {
	_, err = l.svcCtx.Pms.SkuAdd(l.ctx, &pmsclient.SkuAddReq{
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
	return &types.AddSkuResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
