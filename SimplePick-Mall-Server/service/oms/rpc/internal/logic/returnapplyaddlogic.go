package logic

import (
	"SimplePick-Mall-Server/service/oms/model"
	"context"

	"SimplePick-Mall-Server/service/oms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/oms/rpc/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyAddLogic {
	return &ReturnApplyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加退货
func (l *ReturnApplyAddLogic) ReturnApplyAdd(in *oms.ReturnApplyAddReq) (*oms.ReturnApplyAddResp, error) {
	role := &model.ReturnApply{
		UserID:         in.UserID,
		OrderID:        in.OrderID,
		ReturnReasonID: in.ReturnReasonID,
		Status:         in.Status,
		Description:    in.Description,
		ProofPics:      in.ProofPics,
		ReturnAmount:   in.ReturnAmount,
		PlaceId:        in.PlaceId,
		MerchantID:     in.MerchantID,
	}
	_, err := l.svcCtx.ReturnApplyModel.AddReturnApply(role)
	if err != nil {
		//return nil, errors.New("添加失败")
		return nil, err
	}

	return &oms.ReturnApplyAddResp{}, nil
}
