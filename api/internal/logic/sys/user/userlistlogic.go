package user

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.ListUserReq) (*types.ListUserResp, error) {
	resp, err := l.svcCtx.Sys.UserList(l.ctx, &sysclient.UserListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Username: req.Username,
		Status:   req.Status,
		Gender:   req.Gander,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.User
	for _, item := range resp.List {
		listUserData := types.User{
			ID:       item.Id,
			Username: item.Username,
			NickName: item.Nickname,
			Phone:    item.Phone,
			Gander:   item.Gender,
			Avatar:   item.Avatar,
			Email:    item.Email,
			Status:   item.Status,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
			CreateBy: item.CreateBy,
			UpdateBy: item.UpdateBy,
			RoleName: item.RoleName,
		}
		list = append(list, &listUserData)
	}
	return &types.ListUserResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
