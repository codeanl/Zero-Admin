package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sys/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err != nil {
		return nil, err
	}
	var list []*sys.LoginLogListData
	for _, loginlog := range all {
		user, _ := l.svcCtx.UserModel.GetUserByID(loginlog.UserID)
		list = append(list, &sys.LoginLogListData{
			ID:         int64(loginlog.ID),
			Status:     loginlog.Status,
			IP:         loginlog.IP,
			CreateTime: loginlog.CreatedAt.Format("2006-01-02 15:04:05"),
			UserID:     loginlog.UserID,
			Avatar:     user.Avatar,
			Username:   user.Username,
		})
	}
	return &sys.LoginLogListResp{
		Total: total,
		List:  list,
	}, nil

}
