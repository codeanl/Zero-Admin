package MerchantsApply

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyListLogic {
	return &MerchantsApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyListLogic) MerchantsApplyList(req *types.ListMerchantsApplyReq) (*types.ListMerchantsApplyResp, error) {
	resp, err := l.svcCtx.Pms.MerchantsApplyList(l.ctx, &pmsclient.MerchantsApplyListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
		Type:     req.Type,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListMerchantsApplyData
	for _, item := range resp.List {
		listUserData := types.ListMerchantsApplyData{
			ID:             item.ID,
			PrincipalName:  item.PrincipalName,
			PrincipalPhone: item.PrincipalPhone,
			IDCardFront:    item.IDCardFront,
			IDCardReverse:  item.IDCardReverse,
			Name:           item.Name,
			Address:        item.Address,
			Pic:            item.Pic,
			Type:           item.Type,
			Status:         item.Status,
			Approver:       item.Approver,
			ApprovalTime:   item.ApprovalTime,
			Remarks:        item.Remarks,
		}
		list = append(list, listUserData)
	}

	return &types.ListMerchantsApplyResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
