package SubjectProduct

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductDeleteLogic {
	return &SubjectProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectProductDeleteLogic) SubjectProductDelete(req *types.DeleteSubjectProductReq) (resp *types.DeleteSubjectProductResp, err error) {
	_, err = l.svcCtx.Sms.SubjectProductDelete(l.ctx, &smsclient.SubjectProductDeleteAddReq{
		IDs: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeleteSubjectProductResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
