package logic

import (
	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录日志列表
func (l *LoginLogListLogic) LoginLogList(in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	all, total, err := l.svcCtx.LoginLogModel.GetLoginLogList(in)
	log.Println(all[0].UserID)
	if err != nil {
		return nil, err
	}
	var list []*sys.LoginLogListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreateTime = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return &sys.LoginLogListResp{
		Total: total,
		List:  list,
	}, nil
}
