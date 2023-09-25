package Merchants

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsInfoLogic {
	return &MerchantsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsInfoLogic) MerchantsInfo(req *types.MerchantsInfoReq) (*types.MerchantsInfoResp, error) {
	resp, err := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{
		ID:     req.ID,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新失败")
	}
	data := types.ListMerchantsData{
		ID:        resp.ID,
		Name:      resp.Name,
		Principal: resp.Principal,
		Phone:     resp.Phone,
		Address:   resp.Address,
		Pic:       resp.Pic,
		UserID:    resp.UserID,
	}
	return &types.MerchantsInfoResp{
		Code:    200,
		Message: "成功",
		Data:    data,
	}, nil
}
