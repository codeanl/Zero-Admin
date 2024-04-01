package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceDeleteLogic {
	return &PlaceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除自提点
func (l *PlaceDeleteLogic) PlaceDelete(in *pms.PlaceDeleteReq) (*pms.PlaceDeleteResp, error) {
	err := l.svcCtx.PlaceModel.DeletePlaceByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &pms.PlaceDeleteResp{}, nil
}
