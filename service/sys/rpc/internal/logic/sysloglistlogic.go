package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 日志列表
func (l *SysLogListLogic) SysLogList(in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	all, total, err := l.svcCtx.LogModel.GetLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.SysLogListData
	for _, log := range all {
		userinfo, _ := l.svcCtx.UserModel.GetUserByID(log.UserID)
		list = append(list, &sys.SysLogListData{
			Id:         int64(log.ID),
			UserId:     log.UserID,
			Method:     log.Method,
			Operation:  log.Operation,
			Params:     log.Params,
			Time:       log.Time,
			CreateTime: log.CreatedAt.Format("2006-01-02 15:04:05"),
			Ip:         log.IP,
			Username:   userinfo.Username,
			Avatar:     userinfo.Avatar,
		})
	}
	return &sys.SysLogListResp{
		Total: total,
		List:  list,
	}, nil
}
