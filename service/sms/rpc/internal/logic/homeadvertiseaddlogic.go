package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加广告
func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(in *sms.HomeAdvertiseAddReq) (*sms.HomeAdvertiseAddResp, error) {
	role := &model.HomeAdvertise{
		Name:       in.Name,
		Pic:        in.Pic,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		Url:        in.Url,
		Note:       in.Note,
	}
	err := l.svcCtx.HomeAdvertiseModel.AddHomeAdvertise(role)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &sms.HomeAdvertiseAddResp{}, nil
}
