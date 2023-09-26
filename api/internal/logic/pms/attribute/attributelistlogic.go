package attribute

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

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
		Current:             req.Current,
		PageSize:            req.PageSize,
		Name:                req.Name,
		Type:                req.Type,
		AttributeCategoryID: req.AttributeCategoryID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if ii == "商家" {
			isSJ = true
		}
	}
	var list []types.ListAttributeData
	for _, item := range resp.List {
		if isSJ {
			if item.MerchantID == merchant.ID {
				listUserData := types.ListAttributeData{
					Id:                  item.Id,
					AttributeCategoryID: item.AttributeCategoryID,
					Name:                item.Name,
					Type:                item.Type,
					Value:               item.Value,
					Sort:                item.Sort,
				}
				list = append(list, listUserData)
			}
		} else {
			listUserData := types.ListAttributeData{
				Id:                  item.Id,
				AttributeCategoryID: item.AttributeCategoryID,
				Name:                item.Name,
				Type:                item.Type,
				Value:               item.Value,
				Sort:                item.Sort,
			}
			list = append(list, listUserData)
		}
	}
	return &types.ListAttributeResp{
		Code:    200,
		Message: "查询列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
