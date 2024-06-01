package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyAddLogic {
	return &MerchantsApplyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商家申请
func (l *MerchantsApplyAddLogic) MerchantsApplyAdd(in *pms.MerchantsApplyAddReq) (*pms.MerchantsApplyAddResp, error) {
	info := &model.MerchantsApply{
		PrincipalName:  in.PrincipalName,
		PrincipalPhone: in.PrincipalPhone,
		IDCardFront:    in.IDCardFront,
		IDCardReverse:  in.IDCardReverse,
		Name:           in.Name,
		Address:        in.Address,
		Pic:            in.Pic,
		Type:           in.Type,
		Status:         in.Status,
		Remarks:        in.Remarks,
	}
	dd, err := l.svcCtx.MerchantsApplyModel.AddMerchantsApply(info)
	if err != nil {
		return nil, err
	}
	return &pms.MerchantsApplyAddResp{
		ID:             int64(dd.ID),
		PrincipalName:  dd.PrincipalName,
		PrincipalPhone: dd.PrincipalPhone,
		IDCardFront:    dd.IDCardFront,
		IDCardReverse:  dd.IDCardReverse,
		Name:           dd.Name,
		Address:        dd.Address,
		Type:           dd.Type,
		Pic:            dd.Pic,
		Status:         dd.Status,
		Remarks:        dd.Remarks,
	}, nil
}
