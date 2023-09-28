package place

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceInfoLogic {
	return &PlaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceInfoLogic) PlaceInfo(req *types.PlaceInfoReq) (*types.PlaceInfoResp, error) {
	resp, err := l.svcCtx.Sys.PlaceInfo(l.ctx, &sysclient.PlaceInfoReq{
		Id:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return &types.PlaceInfoResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	data := types.ListPlaceData{
		Id:        resp.PlaceInfo.Id,
		Name:      resp.PlaceInfo.Name,
		Place:     resp.PlaceInfo.Place,
		Status:    resp.PlaceInfo.Status,
		Pic:       resp.PlaceInfo.Pic,
		Phone:     resp.PlaceInfo.Phone,
		Principal: resp.PlaceInfo.Principal,
		UserID:    resp.PlaceInfo.UserID,
	}
	return &types.PlaceInfoResp{
		Code:    200,
		Message: "查询成功",
		Data:    data,
	}, nil
}
