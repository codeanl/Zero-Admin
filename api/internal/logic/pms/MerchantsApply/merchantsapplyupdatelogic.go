package MerchantsApply

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyUpdateLogic {
	return &MerchantsApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyUpdateLogic) MerchantsApplyUpdate(req *types.UpdateMerchantsApplyReq) (resp *types.UpdateMerchantsApplyResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsApplyUpdate(l.ctx, &pmsclient.MerchantsApplyUpdateReq{
		ID:             req.ID,
		PrincipalName:  req.PrincipalName,
		PrincipalPhone: req.PrincipalPhone,
		IDCardFront:    req.IDCardFront,
		IDCardReverse:  req.IDCardReverse,
		Name:           req.Name,
		Address:        req.Address,
		Pic:            req.Pic,
		Status:         req.Status,
		Type:           req.Type,
		Approver:       req.Approver,
		ApprovalTime:   req.ApprovalTime,
		Remarks:        req.Remarks,
	})
	if err != nil {
		return &types.UpdateMerchantsApplyResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateMerchantsApplyResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
