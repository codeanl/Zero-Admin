package sku

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuDeleteLogic {
	return &SkuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuDeleteLogic) SkuDelete(req *types.DeleteSkuReq) (resp *types.DeleteSkuResp, err error) {
	_, err = l.svcCtx.Pms.SkuDelete(l.ctx, &pmsclient.SkuDeleteReq{
		Ids: req.IDs,
	})
	if err != nil {
		return &types.DeleteSkuResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteSkuResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
