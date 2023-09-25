package MerchantsApply

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyAddLogic {
	return &MerchantsApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyAddLogic) MerchantsApplyAdd(req *types.AddMerchantsApplyReq) (*types.AddMerchantsApplyResp, error) {
	resp, err := l.svcCtx.Pms.MerchantsApplyAdd(l.ctx, &pmsclient.MerchantsApplyAddReq{
		PrincipalName:  req.PrincipalName,
		PrincipalPhone: req.PrincipalPhone,
		IDCardFront:    req.IDCardFront,
		IDCardReverse:  req.IDCardReverse,
		Name:           req.Name,
		Address:        req.Address,
		Pic:            req.Pic,
		Type:           req.Type,
		Status:         "0",
		Remarks:        req.Remarks,
	})
	data := types.ListMerchantsApplyData{
		ID:             resp.ID,
		PrincipalName:  resp.PrincipalName,
		PrincipalPhone: resp.PrincipalPhone,
		IDCardFront:    resp.IDCardFront,
		IDCardReverse:  resp.IDCardReverse,
		Name:           resp.Name,
		Address:        resp.Address,
		Pic:            resp.Pic,
		Type:           req.Type,
		Status:         resp.Status,
		Remarks:        resp.Remarks,
	}
	if err != nil {
		return nil, err
	}
	return &types.AddMerchantsApplyResp{
		Code:    200,
		Message: "添加成功",
		Data:    data,
	}, nil
}
