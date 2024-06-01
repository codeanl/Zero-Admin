package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryAddLogic {
	return &CategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加分类
func (l *CategoryAddLogic) CategoryAdd(in *pms.CategoryAddReq) (*pms.CategoryAddResp, error) {
	role := &model.Category{
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
	}
	err := l.svcCtx.CategoryModel.AddCategory(role)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &pms.CategoryAddResp{}, nil
}
