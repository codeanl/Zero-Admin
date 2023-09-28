package HotRecommend

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotRecommendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendDeleteLogic {
	return &HotRecommendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotRecommendDeleteLogic) HotRecommendDelete(req *types.DeleteHotRecommendReq) (resp *types.DeleteHotRecommendResp, err error) {
	_, err = l.svcCtx.Sms.HotRecommendDelete(l.ctx, &smsclient.HotRecommendDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteHotRecommendResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteHotRecommendResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
