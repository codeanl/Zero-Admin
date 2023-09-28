package attribute

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeUpdateLogic {
	return &AttributeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeUpdateLogic) AttributeUpdate(req *types.UpdateAttributeReq) (resp *types.UpdateAttributeResp, err error) {
	_, err = l.svcCtx.Pms.AttributeUpdate(l.ctx, &pmsclient.AttributeUpdateReq{
		Id:                  req.Id,
		AttributeCategoryID: req.AttributeCategoryID,
		Name:                req.Name,
		Type:                req.Type,
		Value:               req.Value,
		Sort:                req.Sort,
	})
	if err != nil {
		return &types.UpdateAttributeResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateAttributeResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
