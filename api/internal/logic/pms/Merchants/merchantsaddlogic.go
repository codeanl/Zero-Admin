package Merchants

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsAddLogic {
	return &MerchantsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsAddLogic) MerchantsAdd(req *types.AddMerchantsReq) (resp *types.AddMerchantsResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsAdd(l.ctx, &pmsclient.MerchantsAddReq{
		Name:      req.Name,
		Principal: req.Principal,
		Phone:     req.Phone,
		Address:   req.Address,
		Pic:       req.Pic,
		UserID:    req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddMerchantsResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
