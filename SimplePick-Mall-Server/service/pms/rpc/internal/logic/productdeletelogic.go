package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *ProductDeleteLogic) ProductDelete(in *pms.ProductDeleteReq) (*pms.ProductDeleteResp, error) {
	err := l.svcCtx.ProductModel.DeleteProductByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.ProductDeleteResp{}, nil

}
