package category

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryDeleteLogic) CategoryDelete(req *types.DeleteCategoryReq) (*types.DeleteCategoryResp, error) {
	_, err := l.svcCtx.Pms.CategoryDelete(l.ctx, &pmsclient.CategoryDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteCategoryResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteCategoryResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
