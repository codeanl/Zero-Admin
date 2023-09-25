package Merchants

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsUpdateLogic {
	return &MerchantsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsUpdateLogic) MerchantsUpdate(req *types.UpdateMerchantsReq) (resp *types.UpdateMerchantsResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsUpdate(l.ctx, &pmsclient.MerchantsUpdateReq{
		ID:        req.ID,
		Name:      req.Name,
		Principal: req.Principal,
		Phone:     req.Phone,
		Address:   req.Address,
		Pic:       req.Pic,
		UserID:    req.UserID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新失败")
	}
	return &types.UpdateMerchantsResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
