package user

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	resp, err := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	if err != nil {
		return &types.UserInfoResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	UserInfo := types.UserInfo{
		ID:       resp.UserInfo.ID,
		Username: resp.UserInfo.Username,
		NickName: resp.UserInfo.Nickname,
		Phone:    resp.UserInfo.Phone,
		Gander:   resp.UserInfo.Gender,
		Avatar:   resp.UserInfo.Avatar,
		Email:    resp.UserInfo.Email,
		Status:   resp.UserInfo.Status,
		CreateAt: resp.UserInfo.CreateAt,
		UpdateAt: resp.UserInfo.UpdateAt,
		CreateBy: resp.UserInfo.CreateBy,
		UpdateBy: resp.UserInfo.UpdateBy,
	}
	//
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	for index, d := range resp.Roles {
		if d == "自提点管理员" {
			place, _ := l.svcCtx.Sys.PlaceInfo(l.ctx, &sysclient.PlaceInfoReq{UserID: userInfo.UserInfo.ID})
			resp.Roles[index] = "自提点管理员（" + place.PlaceInfo.Name + ")"
		}
		if d == "商家" {
			merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
			resp.Roles[index] = "商家（" + merchant.Name + ")"
		}
	}
	//
	data := types.Data{
		UserInfo: UserInfo,
		Routes:   resp.Routes,
		Roles:    resp.Roles,
		Buttons:  resp.Buttons,
	}
	return &types.UserInfoResp{
		Code:    200,
		Data:    data,
		Message: "查询成功",
	}, nil
}
