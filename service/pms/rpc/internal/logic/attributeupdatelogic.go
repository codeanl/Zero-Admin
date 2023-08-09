package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

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
	// todo: add your logic here and delete this line

	return &pms.AttributeUpdateResp{}, nil
}
