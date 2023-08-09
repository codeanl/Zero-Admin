package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pms.AttributeDeleteResp{}, nil
}
