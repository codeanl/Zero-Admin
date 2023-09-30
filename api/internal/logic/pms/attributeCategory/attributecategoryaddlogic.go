package attributeCategory

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
	//
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	userInfo, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	merchant, _ := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{UserID: userInfo.UserInfo.ID})
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
	_, err = l.svcCtx.Pms.AttributeCategoryAdd(l.ctx, &pmsclient.AttributeCategoryAddReq{
		Name:       req.Name,
		ParentID:   req.ParentID,
		MerchantID: req.MerchantID,
	})
	if err != nil {
		return &types.AddAttributeCategoryResp{
			Code:    400,
			Message: "添加失败",
		}, nil
	}
	return &types.AddAttributeCategoryResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
