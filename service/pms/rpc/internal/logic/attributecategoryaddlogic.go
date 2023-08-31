package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryAddLogic {
	return &AttributeCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加属性分类
func (l *AttributeCategoryAddLogic) AttributeCategoryAdd(in *pms.AttributeCategoryAddReq) (*pms.AttributeCategoryAddResp, error) {
	role := &model.AttributeCategory{
		Name:     in.Name,
		ParentId: in.ParentID,
	}
	_, err := l.svcCtx.AttributeCategoryModel.AddAttributeCategory(role)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &pms.AttributeCategoryAddResp{}, nil
}
