package place

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceUpdateLogic) PlaceUpdate(req *types.UpdatePlaceReq) (resp *types.UpdatePlaceResp, err error) {
	_, err = l.svcCtx.Pms.PlaceUpdate(l.ctx, &pmsclient.PlaceUpdateReq{
		Id:        req.Id,
		Name:      req.Name,
		Place:     req.Place,
		Status:    req.Status,
		Pic:       req.Pic,
		Phone:     req.Phone,
		Principal: req.Principal,
		UpdateBy:  l.ctx.Value("username").(string),
		UserID:    req.UserID,
	})
	if err != nil {
		return &types.UpdatePlaceResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdatePlaceResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
