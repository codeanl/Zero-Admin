package attributeCategory

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryUpdateLogic {
	return &AttributeCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeCategoryUpdateLogic) AttributeCategoryUpdate(req *types.UpdateAttributeCategoryReq) (resp *types.UpdateAttributeCategoryResp, err error) {
	_, err = l.svcCtx.Pms.AttributeCategoryUpdate(l.ctx, &pmsclient.AttributeCategoryUpdateReq{
		ID:         req.Id,
		Name:       req.Name,
		ParentID:   req.ParentID,
		MerchantID: req.MerchantID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateAttributeCategoryResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
