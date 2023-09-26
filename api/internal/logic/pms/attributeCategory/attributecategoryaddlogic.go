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

type AttributeCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryAddLogic {
	return &AttributeCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttributeCategoryAddLogic) AttributeCategoryAdd(req *types.AddAttributeCategoryReq) (resp *types.AddAttributeCategoryResp, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
	isSJ := false
	for _, ii := range userInfo.Roles {
		if ii == "商家" {
			isSJ = true
		}
	}
	merchantID := int64(0)
	if isSJ {
		merchantID = merchant.ID
	}
	_, err = l.svcCtx.Pms.AttributeCategoryAdd(l.ctx, &pmsclient.AttributeCategoryAddReq{
		Name:       req.Name,
		ParentID:   req.ParentID,
		MerchantID: merchantID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAttributeCategoryResp{
		Code:    200,
		Message: "添加角色成功",
	}, nil
}
