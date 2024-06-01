package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc CategoryFirstList(CategoryFirstListReq) returns(CategoryListResp);
func (l *CategoryUpdateLogic) CategoryUpdate(in *pms.CategoryUpdateReq) (*pms.CategoryUpdateResp, error) {
	err := l.svcCtx.CategoryModel.UpdateCategory(in.Id, &model.Category{
		ParentId:     in.ParentId,
		Name:         in.Name,
		Level:        in.Level,
		ProductCount: in.Sort,
		ProductUnit:  in.ProductUnit,
		NavStatus:    in.ShowStatus,
		ShowStatus:   in.ShowStatus,
		Sort:         in.Sort,
		Icon:         in.Icon,
		Keywords:     in.Keywords,
		Description:  in.Description,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.CategoryUpdateResp{}, nil
}
