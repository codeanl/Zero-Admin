package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 专题列表
func (l *SubjectListLogic) SubjectList(in *sms.SubjectListReq) (*sms.SubjectListResp, error) {
	all, total, err := l.svcCtx.SubjectModel.GetSubjectList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.SubjectListData
	for _, role := range all {
		list = append(list, &sms.SubjectListData{
			ID:     int64(role.ID),
			Name:   role.Name,
			Pic:    role.Pic,
			Status: role.Status,
			Sort:   role.Sort,
		})
	}
	return &sms.SubjectListResp{
		Total: total,
		List:  list,
	}, nil
}
