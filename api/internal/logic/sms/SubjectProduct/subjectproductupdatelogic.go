package SubjectProduct

import (
	"SimplePick-Mall-Server/common/errorx"
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
		logx.Error(err)
		return nil, errorx.NewDefaultError("更新失败")
	}
	return &types.UpdateSubjectProductResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
