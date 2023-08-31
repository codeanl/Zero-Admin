package MemberLoginLog

import (
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogAddLogic {
	return &MemberLoginLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(req *types.AddMemberLoginLogReq) (resp *types.AddMemberLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
