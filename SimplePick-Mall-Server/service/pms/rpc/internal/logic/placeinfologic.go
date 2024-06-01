package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceInfoLogic {
	return &PlaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点详情
func (l *PlaceInfoLogic) PlaceInfo(in *pms.PlaceInfoReq) (*pms.PlaceInfoResp, error) {
	place, _ := l.svcCtx.PlaceModel.GetPlaceByIdOrUserID(in)
	placeInfo := &pms.PlaceListData{
		Id:        int64(place.ID),
		Name:      place.Name,
		Place:     place.Place,
		Status:    place.Status,
		Pic:       place.Pic,
		Phone:     place.Phone,
		Principal: place.Principal,
		CreateBy:  place.CreateBy,
		UpdateBy:  place.UpdateBy,
		UserID:    place.UserID,
	}
	return &pms.PlaceInfoResp{
		PlaceInfo: placeInfo,
	}, nil

}
