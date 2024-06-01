package HotRecommend

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type HotRecommendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotRecommendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotRecommendListLogic {
	return &HotRecommendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotRecommendListLogic) HotRecommendList(req *types.ListHotRecommendReq) (*types.ListHotRecommendResp, error) {
	//resp, err := l.svcCtx.Sms.HotRecommendList(l.ctx, &smsclient.HotRecommendListReq{
	//	Status:   req.Status,
	//	Current:  req.Current,
	//	PageSize: req.PageSize,
	//})
	//if err != nil {
	//	return nil, errorx.NewDefaultError("查询首页品牌失败")
	//}
	//var list []*types.ListHotRecommendData
	//for _, item := range resp.List {
	//	Info, _ := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{ID: item.ProductId})
	//	productInfo := types.ProductInfo{
	//		Id:            (*Info).List[0].Id,
	//		CategoryID:    (*Info).List[0].CategoryID,
	//		Name:          (*Info).List[0].Name,
	//		Pic:           (*Info).List[0].Pic,
	//		ProductSn:     (*Info).List[0].ProductSn,
	//		SubTitle:      (*Info).List[0].SubTitle,
	//		Description:   (*Info).List[0].Description,
	//		Price:         (*Info).List[0].Price,
	//		OriginalPrice: (*Info).List[0].OriginalPrice,
	//		Stock:         (*Info).List[0].Stock,
	//		Unit:          (*Info).List[0].Unit,
	//		Sale:          (*Info).List[0].Sale,
	//	}
	//	list = append(list, &types.ListHotRecommendData{
	//		Id:          item.ID,
	//		ProductId:   item.ProductId,
	//		Status:      item.Status,
	//		Sort:        item.Sort,
	//		ProductInfo: productInfo,
	//	})
	//}
	//return &types.ListHotRecommendResp{
	//	Data:    list,
	//	Total:   resp.Total,
	//	Code:    200,
	//	Message: "查询首页品牌成功",
	//}, nil
	return &types.ListHotRecommendResp{}, nil
}
