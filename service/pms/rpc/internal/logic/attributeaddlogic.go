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
		CategoryID: in.CategoryID,
		Name:       in.Name,
	}
	attribute, err := l.svcCtx.AttributeModel.AddAttribute(info)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	for _, i := range in.AttributeName {
		_ = l.svcCtx.AttributeValueModel.AddAttributeValue(&model.AttributeValue{
			AttributeID: int64(attribute.ID),
			Name:        i,
		})
	}
	return &pms.AttributeAddResp{}, nil
}
