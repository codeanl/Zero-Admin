package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectAddLogic {
	return &SubjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题
func (l *SubjectAddLogic) SubjectAdd(in *sms.SubjectAddReq) (*sms.SubjectAddResp, error) {
	role := &model.Subject{
		Name:   in.Name,
		Pic:    in.Pic,
		Status: in.Status,
		Sort:   in.Sort,
	}
	_, err := l.svcCtx.SubjectModel.AddSubject(role)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &sms.SubjectAddResp{}, nil
}
