package user

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

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
		Email:    req.Email,
	})
	if err != nil {
		return &types.ListUserResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.User
	for _, item := range resp.List {
		//
		userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: item.Id})
		for index, d := range item.RoleName {
			if d == "自提点管理员" {
				place, _ := l.svcCtx.Sys.PlaceInfo(l.ctx, &sysclient.PlaceInfoReq{UserID: userInfo.UserInfo.ID})
				item.RoleName[index] = "自提点管理员（" + place.PlaceInfo.Name + ")"
			}
			if d == "商家" {
				merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
				item.RoleName[index] = "商家（" + merchant.Name + ")"
			}
		}
		//
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
