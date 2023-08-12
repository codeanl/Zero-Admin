package attribute

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeListLogic {
	return &AttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeListLogic) AttributeList(req *types.ListAttributeReq) (*types.ListAttributeResp, error) {
	resp, err := l.svcCtx.Pms.AttributeList(l.ctx, &pmsclient.AttributeListReq{
		Current:    req.Current,
		PageSize:   req.PageSize,
		Name:       req.Name,
		Type:       req.Type,
		CategoryID: req.CategoryId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []types.ListAttributeData
	for _, item := range resp.List {
		var AttributeValue []types.AttributeValue
		for _, i := range item.AttributeValue {
			AttributeValue = append(AttributeValue, types.AttributeValue{
				Id:          i.Id,
				Name:        i.Name,
				AttributeID: i.AttributeID,
			})
		}
		listUserData := types.ListAttributeData{
			Id:             item.Id,
			CategoryId:     item.CategoryID,
			Name:           item.Name,
			Type:           item.Type,
			AttributeValue: AttributeValue,
		}
		list = append(list, listUserData)
	}
	return &types.ListAttributeResp{
		Code:    200,
		Message: "查询列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
