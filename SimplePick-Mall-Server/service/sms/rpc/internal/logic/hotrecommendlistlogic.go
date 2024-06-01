package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendListLogic {
	return &HotRecommendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 推荐列表
func (l *HotRecommendListLogic) HotRecommendList(in *sms.HotRecommendListReq) (*sms.HotRecommendListResp, error) {
	//all, total, err := l.svcCtx.HotRecommendModel.GetHotRecommendList(in)
	//if err != nil {
	//	return nil, err
	//}
	//var list []*sms.HotRecommendList
	//for _, role := range all {
	//	list = append(list, &sms.HotRecommendList{
	//		ID:        int64(role.ID),
	//		ProductId: role.ProductId,
	//		Status:    role.Status,
	//		Sort:      role.Sort,
	//	})
	//}
	//return &sms.HotRecommendListResp{
	//	Total: total,
	//	List:  list,
	//}, nil
	return nil, nil
}
