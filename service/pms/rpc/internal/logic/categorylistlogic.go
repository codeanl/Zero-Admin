package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分类列表
func (l *CategoryListLogic) CategoryList(in *pms.CategoryListReq) (*pms.CategoryListResp, error) {
	all, total, err := l.svcCtx.CategoryModel.GetCategoryList()
	if err != nil {
		return nil, err
	}
	var list []*pms.CategoryListData
	for _, role := range all {
		list = append(list, &pms.CategoryListData{
			Id:           int64(role.ID),
			ParentId:     role.ParentId,
			Name:         role.Name,
			Level:        role.Level,
			ProductCount: role.Sort,
			ProductUnit:  role.ProductUnit,
			NavStatus:    role.ShowStatus,
			ShowStatus:   role.ShowStatus,
			Sort:         role.Sort,
			Icon:         role.Icon,
			Keywords:     role.Keywords,
			Description:  role.Description,
		})
	}
	return &pms.CategoryListResp{
		Total: total,
		List:  list,
	}, nil
}
