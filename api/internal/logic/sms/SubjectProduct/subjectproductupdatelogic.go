package SubjectProduct

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductUpdateLogic {
	return &SubjectProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectProductUpdateLogic) SubjectProductUpdate(req *types.UpdateSubjectProductReq) (resp *types.UpdateSubjectProductResp, err error) {
	_, err = l.svcCtx.Sms.SubjectProductUpdate(l.ctx, &smsclient.SubjectProductUpdateReq{
		ID:        req.ID,
		SubjectID: req.SubjectID,
		ProductID: req.ProductID,
		Status:    req.Status,
		Sort:      req.Sort,
	})
	if err != nil {
		return &types.UpdateSubjectProductResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateSubjectProductResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
