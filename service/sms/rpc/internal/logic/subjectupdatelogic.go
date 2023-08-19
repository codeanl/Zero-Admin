package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectUpdateLogic {
	return &SubjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题
func (l *SubjectUpdateLogic) SubjectUpdate(in *sms.SubjectUpdateReq) (*sms.SubjectUpdateResp, error) {
	err := l.svcCtx.SubjectModel.UpdateSubject(in.ID, &model.Subject{
		Name:   in.Name,
		Pic:    in.Pic,
		Status: in.Status,
		Sort:   in.Sort,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &sms.SubjectUpdateResp{}, nil
}
