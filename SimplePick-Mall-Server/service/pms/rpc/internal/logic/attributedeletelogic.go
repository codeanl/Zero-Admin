package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeDeleteLogic {
	return &AttributeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除属性
func (l *AttributeDeleteLogic) AttributeDelete(in *pms.AttributeDeleteReq) (*pms.AttributeDeleteResp, error) {
	err := l.svcCtx.AttributeModel.DeleteAttributeByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.AttributeDeleteResp{}, nil
}
