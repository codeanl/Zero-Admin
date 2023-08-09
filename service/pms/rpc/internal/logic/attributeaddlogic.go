package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

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
	// todo: add your logic here and delete this line

	return &pms.AttributeAddResp{}, nil
}
