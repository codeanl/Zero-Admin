package logic

import (
	"context"
	"log"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

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
func (l *PlaceInfoLogic) PlaceInfo(in *sys.PlaceInfoReq) (*sys.PlaceInfoResp, error) {
	place, _ := l.svcCtx.PlaceModel.GetPlaceByIdOrUserID(in)
	placeInfo := &sys.PlaceListData{
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
	log.Print(place)
	return &sys.PlaceInfoResp{
		PlaceInfo: placeInfo,
	}, nil
}
