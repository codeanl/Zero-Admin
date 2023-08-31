package MemberLoginLog

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(req *types.DeleteMemberLoginLogReq) (resp *types.DeleteMemberLoginLogResp, err error) {
	_, err = l.svcCtx.Ums.MemberLoginLogDelete(l.ctx, &umsclient.MemberLoginLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeleteMemberLoginLogResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
