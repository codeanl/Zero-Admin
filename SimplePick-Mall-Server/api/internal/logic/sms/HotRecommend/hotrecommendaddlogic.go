package HotRecommend

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotRecommendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendAddLogic {
	return &HotRecommendAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotRecommendAddLogic) HotRecommendAdd(req *types.AddHotRecommendReq) (resp *types.AddHotRecommendResp, err error) {
	_, err = l.svcCtx.Sms.HotRecommendAdd(l.ctx, &smsclient.HotRecommendAddReq{
		ProductId: req.ProductIds,
	})
	if err != nil {
		return &types.AddHotRecommendResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddHotRecommendResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
