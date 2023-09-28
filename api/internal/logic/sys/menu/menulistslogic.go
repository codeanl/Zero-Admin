package menu

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListsLogic {
	return &MenuListsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListsLogic) MenuLists(req *types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{})
	if err != nil {
		return &types.ListMenuResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	menuItems := make([]types.ListMenuData, 0)
	for _, item := range resp.MenuList {
		menuItem := types.ListMenuData{
			Id:       item.ID,
			Name:     item.Name,
			ParentId: item.ParentID,
			Url:      item.Url,
			Remark:   item.Remark,
			Type:     item.Type,
			Icon:     item.Icon,
			OrderNum: item.OrderNum,
			TAG:      item.TAG,
		}
		menuItems = append(menuItems, menuItem)
	}

	menuTree := buildMenuTree(menuItems, 0)

	return &types.ListMenuResp{
		Code:    200,
		Data:    menuTree,
		Message: "查询成功",
	}, nil
}
func buildMenuTree(menuItems []types.ListMenuData, parentID int64) []types.ListMenuData {
	tree := make([]types.ListMenuData, 0)
	for _, item := range menuItems {
		if item.ParentId == parentID {
			children := buildMenuTree(menuItems, item.Id)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
