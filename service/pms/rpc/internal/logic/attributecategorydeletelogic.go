package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryDeleteLogic {
	return &AttributeCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除属性分类
func (l *AttributeCategoryDeleteLogic) AttributeCategoryDelete(in *pms.AttributeCategoryDeleteAddReq) (*pms.AttributeCategoryDeleteResp, error) {
	err := l.svcCtx.AttributeCategoryModel.DeleteAttributeCategoryByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.AttributeCategoryDeleteResp{}, nil
}
