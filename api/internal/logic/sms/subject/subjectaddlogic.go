package subject

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectAddLogic {
	return &SubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectAddLogic) SubjectAdd(req *types.AddSubjectReq) (resp *types.AddSubjectResp, err error) {
	_, err = l.svcCtx.Sms.SubjectAdd(l.ctx, &smsclient.SubjectAddReq{
		Name:   req.Name,
		Pic:    req.Pic,
		Status: req.Status,
		Sort:   req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddSubjectResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
