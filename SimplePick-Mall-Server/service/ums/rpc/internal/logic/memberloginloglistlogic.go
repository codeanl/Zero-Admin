package logic

import (
	"context"

	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录日志列表
func (l *MemberLoginLogListLogic) MemberLoginLogList(in *ums.MemberLoginLogListReq) (*ums.MemberLoginLogListResp, error) {
	all, total, err := l.svcCtx.MemberLoginLogModel.GetMemberLoginLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*ums.MemberLoginLogListData
	for _, user := range all {
		list = append(list, &ums.MemberLoginLogListData{
			Id:         int64(user.ID),
			UserID:     user.UserID,
			IP:         user.IP,
			CreateTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &ums.MemberLoginLogListResp{
		Total: total,
		List:  list,
	}, nil
}
