package index

import (
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListByCateIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListByCateIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListByCateIDLogic {
	return &ProductListByCateIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListByCateIDLogic) ProductListByCateID(req *types.ListProductByCateIDReq) (resp *types.ListProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
