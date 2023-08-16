package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuListLogic {
	return &SkuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Sku列表
func (l *SkuListLogic) SkuList(in *pms.SkuListReq) (*pms.SkuListResp, error) {
	all, total, err := l.svcCtx.SkuModel.GetSkuList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.SkuListData
	for _, role := range all {
		list = append(list, &pms.SkuListData{
			ID:          int64(role.ID),
			ProductID:   role.ProductID,
			Name:        role.Name,
			Pic:         role.Pic,
			SkuSn:       role.SkuSn,
			Description: role.Description,
			Price:       role.Price,
			Stock:       role.Stock,
			Tag:         role.Tag,
		})
	}
	return &pms.SkuListResp{
		Total: total,
		List:  list,
	}, nil
}
