package loginLog

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"
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
	for _, i := range resp.List {
		user, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.UserInfoReq{Id: i.UserID})
		i.Avatar = user.UserInfo.Avatar
		i.Status = user.UserInfo.Status
		i.Username = user.UserInfo.Username
	}
	if err != nil {
		return &types.ListLoginLogResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListLoginLogData
	jsonData, err := json.Marshal(resp.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListLoginLogResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil

}
