package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsInfoLogic {
	return &MerchantsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家详情
func (l *MerchantsInfoLogic) MerchantsInfo(in *pms.MerchantsInfoReq) (*pms.MerchantsInfoResp, error) {
	info, err := l.svcCtx.MerchantsModel.GetMerchantsByIDOrUserID(in)
	if err != nil {
		return nil, err
	}
	return &pms.MerchantsInfoResp{
		ID:        int64(info.ID),
		Name:      info.Name,
		Principal: info.Principal,
		Phone:     info.Phone,
		Address:   info.Address,
		Pic:       info.Pic,
		UserID:    info.UserID,
	}, nil
}
