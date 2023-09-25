package Merchants

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsListLogic {
	return &MerchantsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsListLogic) MerchantsList(req *types.ListMerchantsReq) (*types.ListMerchantsResp, error) {
	resp, err := l.svcCtx.Pms.MerchantsList(l.ctx, &pmsclient.MerchantsListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Principal: req.Principal,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListMerchantsData
	for _, item := range resp.List {
		listUserData := types.ListMerchantsData{
			ID:        item.ID,
			Name:      item.Name,
			Principal: item.Principal,
			Phone:     item.Phone,
			Address:   item.Address,
			Pic:       item.Pic,
			UserID:    item.UserID,
		}
		list = append(list, listUserData)
	}

	return &types.ListMerchantsResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
