package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeUpdateLogic {
	return &AttributeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新属性
func (l *AttributeUpdateLogic) AttributeUpdate(in *pms.AttributeUpdateReq) (*pms.AttributeUpdateResp, error) {
	err := l.svcCtx.AttributeModel.UpdateAttribute(in.Id, &model.Attribute{
		AttributeCategoryID: in.AttributeCategoryID,
		Name:                in.Name,
		Type:                in.Type,
		Value:               in.Value,
		Sort:                in.Sort,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.AttributeUpdateResp{}, nil
}
