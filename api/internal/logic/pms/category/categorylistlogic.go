package category

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.ListCategoryReq) (*types.ListCategoryResp, error) {
	resp, err := l.svcCtx.Pms.CategoryList(l.ctx, &pmsclient.CategoryListReq{})
	if err != nil {
		return &types.ListCategoryResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	list := make([]types.ListCategoryData, 0)
	for _, item := range resp.List {
		pms, _ := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{CategoryID: item.Id})
		listUserData := types.ListCategoryData{
			Id:          item.Id,
			ParentId:    item.ParentId,
			Name:        item.Name,
			Level:       item.Level,
			ProductUnit: item.ProductUnit,
			NavStatus:   item.ShowStatus,
			ShowStatus:  item.ShowStatus,
			Sort:        item.Sort,
			Icon:        item.Icon,
			Keywords:    item.Keywords,
			Description: item.Description,
			Count:       pms.Total,
		}
		list = append(list, listUserData)
	}
	list = buildTree(list, 0)

	return &types.ListCategoryResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
func buildTree(Items []types.ListCategoryData, parentID int64) []types.ListCategoryData {
	tree := make([]types.ListCategoryData, 0)
	for _, item := range Items {
		if item.ParentId == parentID {
			children := buildTree(Items, item.Id)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
