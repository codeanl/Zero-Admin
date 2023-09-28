package sku

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuListLogic {
	return &SkuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuListLogic) SkuList(req *types.ListSkuReq) (*types.ListSkuResp, error) {
	resp, err := l.svcCtx.Pms.SkuList(l.ctx, &pmsclient.SkuListReq{
		ProductID: req.ProductID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListSkuData
	for _, item := range resp.List {
		listUserData := types.ListSkuData{
			ID:          item.ID,
			ProductID:   item.ProductID,
			Name:        item.Name,
			Pic:         item.Pic,
			SkuSn:       item.SkuSn,
			Description: item.Description,
			Price:       item.Price,
			Stock:       item.Stock,
			Tag:         item.Tag,
			Sale:        item.Sale,
		}
		list = append(list, listUserData)
	}
	return &types.ListSkuResp{
		Code:    200,
		Message: "查询列表成功",
		Data:    list,
	}, nil
}
