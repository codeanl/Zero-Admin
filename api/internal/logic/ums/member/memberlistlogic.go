package member

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberListLogic) MemberList(req *types.ListMemberReq) (*types.ListMemberResp, error) {
	resp, err := l.svcCtx.Ums.MemberList(l.ctx, &umsclient.MemberListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Username: req.Username,
		Phone:    req.Phone,
		Status:   req.Status,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListMemberData
	for _, item := range resp.List {
		listUserData := types.ListMemberData{
			Id:        item.Id,
			Username:  item.Username,
			Password:  item.Password,
			Nickname:  item.Nickname,
			Phone:     item.Phone,
			Status:    item.Status,
			Avatar:    item.Avatar,
			Gender:    item.Gender,
			Email:     item.Email,
			City:      item.City,
			Job:       item.Job,
			Signature: item.Signature,
			CreatTIme: item.CreateTime,
		}
		list = append(list, &listUserData)
	}

	return &types.ListMemberResp{
		Code:    200,
		Message: "查询列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
