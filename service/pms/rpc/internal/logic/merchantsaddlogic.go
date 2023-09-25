package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsAddLogic {
	return &MerchantsAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商家
func (l *MerchantsAddLogic) MerchantsAdd(in *pms.MerchantsAddReq) (*pms.MerchantsAddResp, error) {
	_, exist, _ := l.svcCtx.MerchantsModel.GetMerchantsByPhone(in.Phone)
	if exist {
		return nil, errors.New("商家已存在")
	}
	info := &model.Merchants{
		Name:      in.Name,
		Principal: in.Principal,
		Phone:     in.Phone,
		Address:   in.Address,
		Pic:       in.Pic,
		UserID:    in.UserID,
	}
	_, err := l.svcCtx.MerchantsModel.AddMerchants(info)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &pms.MerchantsAddResp{}, nil
}
