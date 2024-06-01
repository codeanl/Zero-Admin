package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuDeleteLogic {
	return &SkuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除Sku
func (l *SkuDeleteLogic) SkuDelete(in *pms.SkuDeleteReq) (*pms.SkuDeleteResp, error) {
	err := l.svcCtx.SkuModel.DeleteSkuByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.SkuDeleteResp{}, nil
}
