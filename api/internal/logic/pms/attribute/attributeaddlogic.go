package attribute

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"
	"strings"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeAddLogic {
	return &AttributeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeAddLogic) AttributeAdd(req *types.AddAttributeReq) (resp *types.AddAttributeResp, err error) {
	//
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
	//
	_, err = l.svcCtx.Pms.AttributeAdd(l.ctx, &pmsclient.AttributeAddReq{
		AttributeCategoryID: req.AttributeCategoryID,
		Name:                req.Name,
		Type:                req.Type,
		Value:               req.Value,
		Sort:                req.Sort,
		MerchantID:          req.MerchantID,
	})
	if err != nil {
		return &types.AddAttributeResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddAttributeResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
