package attribute

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"
	"strings"

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
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.UserInfoReq{Id: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.Id})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if strings.Contains("商家", ii) {
			isSJ = true
		}
	}
	if isSJ == true {
		req.MerchantID = merchant.ID
	}
	resp, err := l.svcCtx.Pms.AttributeList(l.ctx, &pmsclient.AttributeListReq{
		Current:             req.Current,
		PageSize:            req.PageSize,
		Name:                req.Name,
		Type:                req.Type,
		AttributeCategoryID: req.AttributeCategoryID,
		MerchantID:          req.MerchantID,
	})
	if err != nil {
		return &types.ListAttributeResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []types.ListAttributeData
	for _, item := range resp.List {
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
	return &types.ListAttributeResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
