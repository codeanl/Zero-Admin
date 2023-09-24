package category

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryAddLogic {
	return &CategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryAddLogic) CategoryAdd(req *types.AddCategoryReq) (resp *types.AddCategoryResp, err error) {
	_, err = l.svcCtx.Pms.CategoryAdd(l.ctx, &pmsclient.CategoryAddReq{
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: req.Sort,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.ShowStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddCategoryResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
