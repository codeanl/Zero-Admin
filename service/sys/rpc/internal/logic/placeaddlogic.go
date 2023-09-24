package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceAddLogic {
	return &PlaceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加自提点
func (l *PlaceAddLogic) PlaceAdd(in *sys.PlaceAddReq) (*sys.PlaceAddResp, error) {
	place := &model.Place{
		Name:      in.Name,
		Place:     in.Place,
		Status:    in.Status,
		Pic:       in.Pic,
		Phone:     in.Phone,
		Principal: in.Principal,
		CreateBy:  in.CreateBy,
		UserID:    in.UserID,
	}
	err := l.svcCtx.PlaceModel.AddPlace(place)
	if err != nil {
		return nil, errors.New("添加自提点失败")
	}
	return &sys.PlaceAddResp{
		ID: int64(place.ID),
	}, nil
}
