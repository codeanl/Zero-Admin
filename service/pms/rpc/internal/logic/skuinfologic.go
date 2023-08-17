package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuInfoLogic {
	return &SkuInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// sku详情
func (l *SkuInfoLogic) SkuInfo(in *pms.SkuInfoReq) (*pms.SkuInfoResp, error) {
	sku, _ := l.svcCtx.SkuModel.GetSkuById(in.ID)
	skuInfo := &pms.SkuListData{
		ID:          int64(sku.ID),
		ProductID:   sku.ProductID,
		Name:        sku.Name,
		Pic:         sku.Pic,
		SkuSn:       sku.SkuSn,
		Description: sku.Description,
		Price:       sku.Price,
		Stock:       sku.Stock,
		Sale:        sku.Sale,
		Tag:         sku.Tag,
	}
	return &pms.SkuInfoResp{
		SkuInfo: skuInfo,
	}, nil
}
