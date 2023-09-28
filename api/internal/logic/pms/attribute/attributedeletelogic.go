package attribute

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeDeleteLogic {
	return &AttributeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeDeleteLogic) AttributeDelete(req *types.DeleteAttributeReq) (*types.DeleteAttributeResp, error) {
	_, err := l.svcCtx.Pms.AttributeDelete(l.ctx, &pmsclient.AttributeDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteAttributeResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteAttributeResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
