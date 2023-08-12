package attribute

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeAddLogic {
	return &AttributeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeAddLogic) AttributeAdd(req *types.AddAttributeReq) (resp *types.AddAttributeResp, err error) {
	var updateValue []*pmsclient.AddValue
	for _, i := range req.AddValue {
		updateValue = append(updateValue, &pmsclient.AddValue{
			Name: i.Name,
		})
	}
	_, err = l.svcCtx.Pms.AttributeAdd(l.ctx, &pmsclient.AttributeAddReq{
		Name:       req.Name,
		CategoryID: req.CategoryId,
		Type:       req.Type,
		AddValue:   updateValue,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAttributeResp{
		Code:    200,
		Message: "添加角色成功",
	}, nil
}