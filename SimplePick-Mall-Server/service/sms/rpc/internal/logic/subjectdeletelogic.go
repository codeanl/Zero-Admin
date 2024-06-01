package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectDeleteLogic {
	return &SubjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题
func (l *SubjectDeleteLogic) SubjectDelete(in *sms.SubjectDeleteAddReq) (*sms.SubjectDeleteResp, error) {
	err := l.svcCtx.SubjectModel.DeleteSubjectByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sms.SubjectDeleteResp{}, nil
}
