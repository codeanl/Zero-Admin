package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单列表
func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	err := l.svcCtx.MenuModel.UpdateMenu(in.Id, &model.Menu{
		Name:     in.Name,
		ParentID: in.ParentId,
		Url:      in.Url,
		Remark:   in.Remark,
		Type:     in.Type,
		Icon:     in.Icon,
		OrderNum: in.OrderNum,
		UpdateBy: in.UpdateBy,
		TAG:      in.TAG,
	})
	if err != nil {
		return nil, errors.New("更新菜单失败")
	}
	return &sys.MenuUpdateResp{}, nil
}
