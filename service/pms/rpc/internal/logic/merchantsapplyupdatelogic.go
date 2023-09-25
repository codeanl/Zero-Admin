package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyUpdateLogic {
	return &MerchantsApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商家申请
func (l *MerchantsApplyUpdateLogic) MerchantsApplyUpdate(in *pms.MerchantsApplyUpdateReq) (*pms.MerchantsApplyUpdateResp, error) {
	err := l.svcCtx.MerchantsApplyModel.UpdateMerchantsApply(in.ID, &model.MerchantsApply{
		PrincipalName:  in.PrincipalName,
		PrincipalPhone: in.PrincipalPhone,
		IDCardFront:    in.IDCardFront,
		IDCardReverse:  in.IDCardReverse,
		Name:           in.Name,
		Address:        in.Address,
		Pic:            in.Pic,
		Type:           in.Type,
		Status:         in.Status,
		Approver:       in.Approver,
		ApprovalTime:   in.ApprovalTime,
		Remarks:        in.Remarks,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.MerchantsApplyUpdateResp{}, nil
}
