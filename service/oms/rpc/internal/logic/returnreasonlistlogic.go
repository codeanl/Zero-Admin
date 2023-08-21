package logic

import (
	"SimplePick-Mall-Server/service/oms/rpc/omsclient"
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnReasonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnReasonListLogic {
	return &ReturnReasonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 退货原因列表
func (l *ReturnReasonListLogic) ReturnReasonList(in *oms.ReturnReasonListReq) (*oms.ReturnReasonListResp, error) {
	order, total, err := l.svcCtx.ReturnReasonModel.GetReturnReasonList(in)
	if err != nil {
		return nil, err
	}
	var List []*omsclient.ReturnReasonListData
	for _, i := range order {
		List = append(List, &omsclient.ReturnReasonListData{
			ID:     int64(i.ID),
			Name:   i.Name,
			Status: i.Status,
		})
	}
	return &oms.ReturnReasonListResp{
		Total: total,
		List:  List,
	}, nil
}
