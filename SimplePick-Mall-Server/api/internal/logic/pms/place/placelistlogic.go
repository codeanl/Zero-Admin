package place

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceListLogic) PlaceList(req *types.ListPlaceReq) (*types.ListPlaceResp, error) {
	resp, err := l.svcCtx.Pms.PlaceList(l.ctx, &pmsclient.PlaceListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Name:      req.Name,
		Place:     req.Place,
		Phone:     req.Phone,
		Principal: req.Principal,
	})
	if err != nil {
		return &types.ListPlaceResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListPlaceData
	for _, item := range resp.List {
		listUserData := types.ListPlaceData{
			Id:        item.Id,
			Name:      item.Name,
			Place:     item.Place,
			Status:    item.Status,
			Pic:       item.Pic,
			Phone:     item.Phone,
			Principal: item.Principal,
			UserID:    item.UserID,
		}
		list = append(list, &listUserData)
	}

	return &types.ListPlaceResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
