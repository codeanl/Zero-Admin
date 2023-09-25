package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsListLogic {
	return &MerchantsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家列表
func (l *MerchantsListLogic) MerchantsList(in *pms.MerchantsListReq) (*pms.MerchantsListResp, error) {
	all, total, err := l.svcCtx.MerchantsModel.GetMerchantsList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.MerchantsListData
	for _, role := range all {
		list = append(list, &pms.MerchantsListData{
			ID:        int64(role.ID),
			Name:      role.Name,
			Principal: role.Principal,
			Phone:     role.Phone,
			Address:   role.Address,
			Pic:       role.Pic,
			UserID:    role.UserID,
		})
	}
	return &pms.MerchantsListResp{
		Total: total,
		List:  list,
	}, nil
}
