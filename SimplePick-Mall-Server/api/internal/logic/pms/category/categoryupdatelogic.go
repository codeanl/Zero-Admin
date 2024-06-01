package category

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(req *types.UpdateCategoryReq) (*types.UpdateCategoryResp, error) {
	_, err := l.svcCtx.Pms.CategoryUpdate(l.ctx, &pmsclient.CategoryUpdateReq{
		Id:           req.Id,
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
		return &types.UpdateCategoryResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateCategoryResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
