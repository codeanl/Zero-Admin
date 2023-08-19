package index

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		return nil, errorx.NewDefaultError("查询失败")
	}
	list := make([]types.ListCategoryData, 0)
	for _, item := range resp.List {
		productList, _ := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{
			CategoryID: item.Id,
			PageNum:    1,
			PageSize:   3,
		})
		var ProductList []types.ListProductData
		for _, i := range productList.List {
			ProductList = append(ProductList, types.ListProductData{
				Id:            i.Id,
				CategoryID:    i.CategoryID,
				Name:          i.Name,
				Pic:           i.Pic,
				ProductSn:     i.ProductSn,
				Desc:          i.Desc,
				OriginalPrice: i.OriginalPrice,
				Unit:          i.Unit,
				Price:         i.Price,
			})
		}
		listData := types.ListCategoryData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.Sort,
			ProductUnit:  item.ProductUnit,
			NavStatus:    item.ShowStatus,
			ShowStatus:   item.ShowStatus,
			Sort:         item.Sort,
			Icon:         item.Icon,
			Keywords:     item.Keywords,
			Description:  item.Description,
			ProductList:  ProductList,
		}
		list = append(list, listData)
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
