package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsUpdateLogic {
	return &MerchantsUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商家
func (l *MerchantsUpdateLogic) MerchantsUpdate(in *pms.MerchantsUpdateReq) (*pms.MerchantsUpdateResp, error) {
	err := l.svcCtx.MerchantsModel.UpdateMerchants(in.ID, &model.Merchants{
		Name:      in.Name,
		Principal: in.Principal,
		Phone:     in.Phone,
		Address:   in.Address,
		Pic:       in.Pic,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.MerchantsUpdateResp{}, nil
}
