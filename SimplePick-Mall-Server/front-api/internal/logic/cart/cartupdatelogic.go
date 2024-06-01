package cart

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartUpdateLogic {
	return &CartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *CartUpdateLogic) CartUpdate(req *types.UpdateCartReq) (resp *types.UpdateCartResp, err error) {
	_, err = l.svcCtx.Oms.CartUpdate(l.ctx, &oms.CartUpdateReq{
		ID:     req.ID,
		UserID: req.UserID,
		SkuID:  req.SkuID,
		Count:  req.Count,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateCartResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
