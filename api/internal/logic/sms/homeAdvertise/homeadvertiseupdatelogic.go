package homeAdvertise

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(req *types.UpdateHomeAdvertiseReq) (resp *types.UpdateHomeAdvertiseResp, err error) {
	_, err = l.svcCtx.Sms.HomeAdvertiseUpdate(l.ctx, &smsclient.HomeAdvertiseUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Pic:    req.Pic,
		Status: req.Status,
		Url:    req.Url,
		Note:   req.Note,
		Sort:   req.Sort,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateHomeAdvertiseResp{
		Code:    200,
		Message: "更新用户成功",
	}, nil
}
