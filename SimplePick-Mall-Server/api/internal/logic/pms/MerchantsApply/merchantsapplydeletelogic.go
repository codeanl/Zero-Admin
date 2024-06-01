package MerchantsApply

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyDeleteLogic {
	return &MerchantsApplyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyDeleteLogic) MerchantsApplyDelete(req *types.DeleteMerchantsApplyReq) (resp *types.DeleteMerchantsApplyResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsApplyDelete(l.ctx, &pmsclient.MerchantsApplyDeleteReq{
		IDs: req.Ids,
	})
	if err != nil {
		return &types.DeleteMerchantsApplyResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteMerchantsApplyResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
