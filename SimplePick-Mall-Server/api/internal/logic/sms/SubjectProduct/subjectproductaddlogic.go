package SubjectProduct

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductAddLogic {
	return &SubjectProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectProductAddLogic) SubjectProductAdd(req *types.AddSubjectProductReq) (resp *types.AddSubjectProductResp, err error) {
	_, err = l.svcCtx.Sms.SubjectProductAdd(l.ctx, &smsclient.SubjectProductAddReq{
		SubjectID: req.SubjectID,
		ProductID: req.ProductID,
		Status:    req.Status,
		Sort:      req.Sort,
	})
	if err != nil {
		return &types.AddSubjectProductResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddSubjectProductResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
