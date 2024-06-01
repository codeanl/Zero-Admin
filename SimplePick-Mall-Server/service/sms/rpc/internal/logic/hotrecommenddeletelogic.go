package logic

import (
	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotRecommendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendDeleteLogic {
	return &HotRecommendDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除推荐
func (l *HotRecommendDeleteLogic) HotRecommendDelete(in *sms.HotRecommendDeleteReq) (*sms.HotRecommendDeleteResp, error) {
	//err := l.svcCtx.HotRecommendModel.DeleteHotRecommendByIds(in.Ids)
	//if err != nil {
	//	return nil, errors.New("删除失败")
	//}
	return &sms.HotRecommendDeleteResp{}, nil
}
