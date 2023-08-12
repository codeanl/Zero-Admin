package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 广告列表
func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	all, total, err := l.svcCtx.HomeAdvertiseModel.GetHomeAdvertiseList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.HomeAdvertiseListData
	for _, role := range all {
		list = append(list, &sms.HomeAdvertiseListData{
			Id:         int64(role.ID),
			Name:       role.Name,
			Pic:        role.Pic,
			Status:     role.Status,
			ClickCount: role.ClickCount,
			Url:        role.Url,
			Note:       role.Note,
		})
	}
	return &sms.HomeAdvertiseListResp{
		Total: total,
		List:  list,
	}, nil
}
