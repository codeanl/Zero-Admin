package attributeCategory

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryDeleteLogic {
	return &AttributeCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeCategoryDeleteLogic) AttributeCategoryDelete(req *types.DeleteAttributeCategoryReq) (resp *types.DeleteAttributeCategoryResp, err error) {
	_, err = l.svcCtx.Pms.AttributeCategoryDelete(l.ctx, &pmsclient.AttributeCategoryDeleteAddReq{
		IDs: req.Ids,
	})
	if err != nil {
		return &types.DeleteAttributeCategoryResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteAttributeCategoryResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
