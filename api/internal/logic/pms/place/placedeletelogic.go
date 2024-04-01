package place

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceDeleteLogic {
	return &PlaceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceDeleteLogic) PlaceDelete(req *types.DeletePlaceReq) (resp *types.DeletePlaceResp, err error) {
	_, err = l.svcCtx.Pms.PlaceDelete(l.ctx, &pmsclient.PlaceDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeletePlaceResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeletePlaceResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
