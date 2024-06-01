package HotRecommend

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotRecommendUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendUpdateLogic {
	return &HotRecommendUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotRecommendUpdateLogic) HotRecommendUpdate(req *types.UpdateHotRecommendReq) (resp *types.UpdateHotRecommendResp, err error) {
	_, err = l.svcCtx.Sms.HotRecommendUpdate(l.ctx, &smsclient.HotRecommendUpdateReq{
		ID:     req.Id,
		Status: req.Status,
		Sort:   req.Sort,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新首页品牌失败")
	}
	return &types.UpdateHotRecommendResp{
		Code:    200,
		Message: "更新首页品牌成功",
	}, nil
}
