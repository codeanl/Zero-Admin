package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotRecommendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendAddLogic {
	return &HotRecommendAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加推荐
func (l *HotRecommendAddLogic) HotRecommendAdd(in *sms.HotRecommendAddReq) (*sms.HotRecommendAddResp, error) {
	for _, i := range in.ProductId {
		_, exist, err := l.svcCtx.HotRecommendModel.ExistHotRecommendById(i)
		if err != nil {
			return nil, err
		}
		//存在的时候不添加
		if !exist {
			err = l.svcCtx.HotRecommendModel.AddHotRecommend(&model.HotRecommend{
				ProductId: i,
				Sort:      1,   //默认
				Status:    "1", //默认
			})
			if err != nil {
				return nil, err
			}
		}
	}
	return &sms.HotRecommendAddResp{}, nil
}
