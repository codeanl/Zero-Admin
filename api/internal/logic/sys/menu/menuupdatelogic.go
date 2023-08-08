package menu

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	_, err = l.svcCtx.Sys.MenuUpdate(l.ctx, &sysclient.MenuUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Remark:   req.Remark,
		Type:     req.Type,
		Icon:     req.Icon,
		OrderNum: req.OrderNum,
		TAG:      req.TAG,
		UpdateBy: l.ctx.Value("username").(string),
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateMenuResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
