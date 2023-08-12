package attribute

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/common/errorx"
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
	var updateValue []*pmsclient.UpdateValue
	for _, i := range req.UpdateValue {
		updateValue = append(updateValue, &pmsclient.UpdateValue{
			Value: i.Name,
		})
	}
	_, err = l.svcCtx.Pms.AttributeUpdate(l.ctx, &pmsclient.AttributeUpdateReq{
		Id:          req.Id,
		Name:        req.Name,
		Type:        req.Type,
		CategoryID:  req.CategoryId,
		UpdateValue: updateValue,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateAttributeResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
