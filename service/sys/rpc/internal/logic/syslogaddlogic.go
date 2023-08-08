package logic

import (
	"SimplePick-Mall-Server/service/sys/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加日志
func (l *SysLogAddLogic) SysLogAdd(in *sys.LogAddReq) (*sys.LogAddResp, error) {
	log := &model.Log{
		UserID:    in.UserId,
		Operation: in.Operation,
		Method:    in.Method,
		Params:    in.Params,
		Time:      in.Time,
		IP:        in.Ip,
	}
	err := l.svcCtx.LogModel.AddLog(log)
	if err != nil {
		return nil, errors.New("添加日志失败")
	}
	return &sys.LogAddResp{}, nil
}
