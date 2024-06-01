package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyListLogic {
	return &MerchantsApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家申请列表
func (l *MerchantsApplyListLogic) MerchantsApplyList(in *pms.MerchantsApplyListReq) (*pms.MerchantsApplyListResp, error) {
	all, total, err := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.MerchantApplysListData
	for _, role := range all {
		list = append(list, &pms.MerchantApplysListData{
			ID:             int64(role.ID),
			PrincipalName:  role.PrincipalName,
			PrincipalPhone: role.PrincipalPhone,
			IDCardFront:    role.IDCardFront,
			IDCardReverse:  role.IDCardReverse,
			Name:           role.Name,
			Address:        role.Address,
			Pic:            role.Pic,
			Type:           role.Type,
			Status:         role.Status,
			Approver:       role.Approver,
			ApprovalTime:   role.ApprovalTime,
			Remarks:        role.Remarks,
		})
	}
	return &pms.MerchantsApplyListResp{
		Total: total,
		List:  list,
	}, nil
}
