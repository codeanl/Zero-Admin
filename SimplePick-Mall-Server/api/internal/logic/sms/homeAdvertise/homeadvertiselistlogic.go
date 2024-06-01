package homeAdvertise

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(req *types.ListHomeAdvertiseReq) (*types.ListHomeAdvertiseResp, error) {
	resp, err := l.svcCtx.Sms.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListHomeAdvertiseResp{
			Code:    200,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListHomeAdvertiseData
	for _, item := range resp.List {
		listUserData := types.ListHomeAdvertiseData{
			Id:     item.Id,
			Name:   item.Name,
			Pic:    item.Pic,
			Status: item.Status,
			Url:    item.Url,
			Note:   item.Note,
			Sort:   item.Sort,
		}
		list = append(list, &listUserData)
	}
	return &types.ListHomeAdvertiseResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
