package systemLog

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSystemLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemLogListLogic {
	return &SystemLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SystemLogListLogic) SystemLogList(req *types.ListSysLogReq) (*types.ListSysLogResp, error) {
	resp, err := l.svcCtx.Sys.SysLogList(l.ctx, &sysclient.SysLogListReq{
		Current:  req.PageNum,
		PageSize: req.PageSize,
		Method:   req.Method,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListSysLogData
	for _, item := range resp.List {
		listUserData := types.ListSysLogData{
			Id:         item.Id,
			UserId:     item.UserId,
			Username:   item.Username,
			Operation:  item.Operation,
			Method:     item.Method,
			Params:     item.Params,
			Time:       item.Time,
			Ip:         item.Ip,
			Avatar:     item.Avatar,
			CreateTime: item.CreateTime,
		}
		list = append(list, &listUserData)
	}
	return &types.ListSysLogResp{
		Code:     200,
		Message:  "查询登录日志列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
