package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

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
func (l *PlaceListLogic) PlaceList(in *pms.PlaceListReq) (*pms.PlaceListResp, error) {
	all, total, err := l.svcCtx.PlaceModel.GetPlaceList(in)
	println(all)
	if err != nil {
		return nil, err
	}
	var list []*pms.PlaceListData
	for _, role := range all {
		list = append(list, &pms.PlaceListData{
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
	return &pms.PlaceListResp{
		Total: total,
		List:  list,
	}, nil

}
