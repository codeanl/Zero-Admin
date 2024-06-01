package logic

import (
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnApplyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyInfoLogic {
	return &ReturnApplyInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 退货详情
func (l *ReturnApplyInfoLogic) ReturnApplyInfo(in *oms.ReturnApplyInfoReq) (*oms.ReturnApplyInfoResp, error) {
	info, _ := l.svcCtx.ReturnApplyModel.GetReturnApplyByOrderID(in.OrderID)
	//if info.ID == 0 {
	//	return nil, nil
	//}
	reason, _ := l.svcCtx.ReturnReasonModel.GetReturnReasonById(info.ReturnReasonID)
	return &oms.ReturnApplyInfoResp{
		ID:               int64(info.ID),
		UserID:           info.UserID,
		OrderID:          info.OrderID,
		ReturnReasonID:   info.ReturnReasonID,
		ReturnReasonName: reason.Name,
		Status:           info.Status,
		Description:      info.Description,
		ProofPics:        info.ProofPics,
		ReturnAmount:     info.ReturnAmount,
		PlaceId:          info.PlaceId,
		MerchantID:       info.MerchantID,
	}, nil

}
