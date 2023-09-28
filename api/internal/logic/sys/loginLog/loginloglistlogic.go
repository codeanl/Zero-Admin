package loginLog

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogListLogic) LoginLogList(req *types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.Sys.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Username: req.Username,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListLoginLogResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListLoginLogData
	for _, item := range resp.List {
		listUserData := types.ListLoginLogData{
			Id:         item.ID,
			UserId:     item.UserID,
			Status:     item.Status,
			Ip:         item.IP,
			CreateTime: item.CreateTime,
			Avatar:     item.Avatar,
			Username:   item.Username,
		}
		list = append(list, &listUserData)
	}
	return &types.ListLoginLogResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil

}
