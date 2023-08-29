package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeAddLogic {
	return &AttributeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加属性
func (l *AttributeAddLogic) AttributeAdd(in *pms.AttributeAddReq) (*pms.AttributeAddResp, error) {
	info := &model.Attribute{
		AttributeCategoryID: in.AttributeCategoryID,
		Name:                in.Name,
		Type:                in.Type,
		Value:               in.Value,
		Sort:                in.Sort,
	}
	_, err := l.svcCtx.AttributeModel.AddAttribute(info)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &pms.AttributeAddResp{}, nil
}
