package MemberLoginLog

import (
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(req *types.ListMemberLoginLogReq) (*types.ListMemberLoginLogResp, error) {
	resp, err := l.svcCtx.Ums.MemberLoginLogList(l.ctx, &umsclient.MemberLoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserID:   req.UserID,
	})
	if err != nil {
		return &types.ListMemberLoginLogResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListMemberLoginLogData
	for _, item := range resp.List {
		listUserData := types.ListMemberLoginLogData{
			Id:         item.Id,
			UserID:     item.UserID,
			IP:         item.IP,
			CreateTime: item.CreateTime,
		}
		list = append(list, &listUserData)
	}
	return &types.ListMemberLoginLogResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
