package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新广告
func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(in *sms.HomeAdvertiseUpdateReq) (*sms.HomeAdvertiseUpdateResp, error) {
	err := l.svcCtx.HomeAdvertiseModel.UpdateHomeAdvertise(in.Id, &model.HomeAdvertise{
		Name:       in.Name,
		Pic:        in.Pic,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		Url:        in.Url,
		Note:       in.Note,
		Sort:       in.Sort,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &sms.HomeAdvertiseUpdateResp{}, nil

}
