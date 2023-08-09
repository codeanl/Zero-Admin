package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeListLogic {
	return &AttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 属性列表
func (l *AttributeListLogic) AttributeList(in *pms.AttributeListReq) (*pms.AttributeListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.AttributeListResp{}, nil
}
