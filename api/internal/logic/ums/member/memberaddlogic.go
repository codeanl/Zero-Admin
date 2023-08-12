package member

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req *types.AddMemberReq) (resp *types.AddMemberResp, err error) {
	_, err = l.svcCtx.Ums.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("添加失败")
	}
	return &types.AddMemberResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
