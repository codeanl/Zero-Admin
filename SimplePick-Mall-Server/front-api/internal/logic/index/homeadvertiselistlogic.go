package index

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		Status: "1",
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListHomeAdvertiseData
	for _, item := range resp.List {
		listUserData := types.ListHomeAdvertiseData{
			Id:     item.Id,
			Name:   item.Name,
			Pic:    item.Pic,
			Status: item.Status,
			Url:    item.Url,
			Note:   item.Note,
		}
		list = append(list, listUserData)
	}
	return &types.ListHomeAdvertiseResp{
		Code:    200,
		Message: "success",
		Data:    list,
		Total:   resp.Total,
	}, nil
}
