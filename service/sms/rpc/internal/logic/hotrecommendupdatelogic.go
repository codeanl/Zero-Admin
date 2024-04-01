package logic

import (
	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotRecommendUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendUpdateLogic {
	return &HotRecommendUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新推荐
func (l *HotRecommendUpdateLogic) HotRecommendUpdate(in *sms.HotRecommendUpdateReq) (*sms.HotRecommendUpdateResp, error) {
	//err := l.svcCtx.HotRecommendModel.UpdateHotRecommend(in.ID, &model.HotRecommend{
	//	Status: in.Status,
	//	Sort:   in.Sort,
	//})
	//if err != nil {
	//	return nil, errors.New("更新失败")
	//}
	return &sms.HotRecommendUpdateResp{}, nil
}
