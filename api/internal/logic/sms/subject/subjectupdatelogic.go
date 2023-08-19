package subject

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectUpdateLogic {
	return &SubjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectUpdateLogic) SubjectUpdate(req *types.UpdateSubjectReq) (resp *types.UpdateSubjectResp, err error) {
	_, err = l.svcCtx.Sms.SubjectUpdate(l.ctx, &smsclient.SubjectUpdateReq{
		ID:     req.ID,
		Name:   req.Name,
		Pic:    req.Pic,
		Status: req.Status,
		Sort:   req.Sort,
	})
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewDefaultError("更新失败")
	}
	return &types.UpdateSubjectResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
