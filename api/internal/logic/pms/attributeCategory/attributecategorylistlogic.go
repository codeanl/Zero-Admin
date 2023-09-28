package attributeCategory

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryListLogic {
	return &AttributeCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeCategoryListLogic) AttributeCategoryList(req *types.ListAttributeCategoryReq) (*types.ListAttributeCategoryResp, error) {
	resp, err := l.svcCtx.Pms.AttributeCategoryList(l.ctx, &pmsclient.AttributeCategoryListReq{})
	if err != nil {
		return &types.ListAttributeCategoryResp{
			Code:    200,
			Message: "查询失败",
		}, nil
	}
	list := make([]types.ListAttributeCategoryData, 0)
	//
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if ii == "商家" {
			isSJ = true
		}
	}
	//
	for _, item := range resp.List {
		listUserData := types.ListAttributeCategoryData{
			Id:         item.ID,
			Name:       item.Name,
			ParentID:   item.ParentID,
			MerchantID: item.MerchantID,
		}
		if isSJ {
			if merchant.ID == item.MerchantID {
				list = append(list, listUserData)
			}
		} else {
			list = append(list, listUserData)
		}

	}
	list = buildTree(list, 0)

	return &types.ListAttributeCategoryResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
func buildTree(Items []types.ListAttributeCategoryData, parentID int64) []types.ListAttributeCategoryData {
	tree := make([]types.ListAttributeCategoryData, 0)
	for _, item := range Items {
		if item.ParentID == parentID {
			children := buildTree(Items, item.Id)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
