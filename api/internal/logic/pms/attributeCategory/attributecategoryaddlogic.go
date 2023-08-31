package attributeCategory

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryAddLogic {
	return &AttributeCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeCategoryAddLogic) AttributeCategoryAdd(req *types.AddAttributeCategoryReq) (resp *types.AddAttributeCategoryResp, err error) {
	_, err = l.svcCtx.Pms.AttributeCategoryAdd(l.ctx, &pmsclient.AttributeCategoryAddReq{
		Name:     req.Name,
		ParentID: req.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAttributeCategoryResp{
		Code:    200,
		Message: "添加角色成功",
	}, nil
}
