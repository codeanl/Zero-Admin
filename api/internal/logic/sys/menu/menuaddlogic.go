package menu

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req *types.AddMenuReq) (resp *types.AddMenuResp, err error) {
	_, err = l.svcCtx.Sys.MenuAdd(l.ctx, &sysclient.MenuAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Remark:   req.Remark,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		CreateBy: l.ctx.Value("username").(string),
		TAG:      req.TAG,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("添加菜单失败")
	}
	return &types.AddMenuResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
