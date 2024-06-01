package logic

import (
	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单列表
func (l *MenuListLogic) MenuList(in *sys.MenuListReq) (*sys.MenuListResp, error) {
	all, _, err := l.svcCtx.MenuModel.GetMenuList()
	if err != nil {
		return nil, err
	}
	var list []*sys.MenuList
	for _, i := range all {
		list = append(list, &sys.MenuList{
			ID:       int64(i.ID),
			Name:     i.Name,
			ParentID: i.ParentID,
			Url:      i.Url,
			Type:     i.Type,
			Icon:     i.Icon,
			OrderNum: i.OrderNum,
			CreateBy: i.CreateBy,
			UpdateBy: i.UpdateBy,
			Remark:   i.Remark,
			TAG:      i.TAG,
		})
	}
	return &sys.MenuListResp{
		MenuList: list,
	}, nil
}
