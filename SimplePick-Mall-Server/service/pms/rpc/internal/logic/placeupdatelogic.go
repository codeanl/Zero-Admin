package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新自提点
func (l *PlaceUpdateLogic) PlaceUpdate(in *pms.PlaceUpdateReq) (*pms.PlaceUpdateResp, error) {
	err := l.svcCtx.PlaceModel.UpdatePlace(in.Id, &model.Place{
		Name:      in.Name,
		Place:     in.Place,
		Status:    in.Status,
		Pic:       in.Pic,
		Phone:     in.Phone,
		Principal: in.Principal,
		UpdateBy:  in.UpdateBy,
		UserID:    in.UserID,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.PlaceUpdateResp{}, nil
}
