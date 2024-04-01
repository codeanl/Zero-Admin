package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加菜单
func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	err := l.svcCtx.MenuModel.AddMenu(&model.Menu{
		Name:     in.Name,
		ParentID: in.ParentId,
		Url:      in.Url,
		Type:     in.Type,
		Icon:     in.Icon,
		OrderNum: in.OrderNum,
		CreateBy: in.CreateBy,
		Remark:   in.Remark,
		TAG:      in.TAG,
	})
	if err != nil {
		return nil, errors.New("添加菜单失败")
	}
	return &sys.MenuAddResp{}, nil
}
