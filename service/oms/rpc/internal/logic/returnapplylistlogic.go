package logic

import (
	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyListLogic {
	return &ReturnApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 退货列表
func (l *ReturnApplyListLogic) ReturnApplyList(in *oms.ReturnApplyListReq) (*oms.ReturnApplyListResp, error) {
	order, total, err := l.svcCtx.ReturnApplyModel.GetReturnApplyList(in)
	if err != nil {
		return nil, err
	}
	var List []*omsclient.ReturnApplyListData
	for _, i := range order {
		reason, _ := l.svcCtx.ReturnReasonModel.GetReturnReasonById(i.ReturnReasonID)
		List = append(List, &omsclient.ReturnApplyListData{
			ID:               int64(i.ID),
			UserID:           i.UserID,
			OrderID:          i.OrderID,
			ReturnReasonID:   i.ReturnReasonID,
			ReturnReasonName: reason.Name,
			Status:           i.Status,
			Description:      i.Description,
			ProofPics:        i.ProofPics,
			ReturnAmount:     i.ReturnAmount,
			PlaceId:          i.PlaceId,
			MerchantID:       i.MerchantID,
		})
	}
	return &oms.ReturnApplyListResp{
		Total: total,
		List:  List,
	}, nil
}
