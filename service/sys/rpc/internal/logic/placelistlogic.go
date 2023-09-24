package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点列表
func (l *PlaceListLogic) PlaceList(in *sys.PlaceListReq) (*sys.PlaceListResp, error) {
	all, total, err := l.svcCtx.PlaceModel.GetPlaceList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.PlaceListData
	for _, role := range all {
		list = append(list, &sys.PlaceListData{
			Id:        int64(role.ID),
			Name:      role.Name,
			Place:     role.Place,
			Status:    role.Status,
			Pic:       role.Pic,
			Phone:     role.Phone,
			Principal: role.Principal,
			UserID:    role.UserID,
		})
	}
	return &sys.PlaceListResp{
		Total: total,
		List:  list,
	}, nil
}
