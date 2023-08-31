package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录日志删除
func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(in *ums.MemberLoginLogDeleteReq) (*ums.MemberLoginLogDeleteResp, error) {
	err := l.svcCtx.MemberModel.DeleteByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}

	return &ums.MemberLoginLogDeleteResp{}, nil
}
