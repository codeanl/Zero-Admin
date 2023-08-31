package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryUpdateLogic {
	return &AttributeCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新属性分类
func (l *AttributeCategoryUpdateLogic) AttributeCategoryUpdate(in *pms.AttributeCategoryUpdateReq) (*pms.AttributeCategoryUpdateResp, error) {
	err := l.svcCtx.AttributeCategoryModel.UpdateAttributeCategory(in.ID, &model.AttributeCategory{
		Name:     in.Name,
		ParentId: in.ParentID,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.AttributeCategoryUpdateResp{}, nil
}
