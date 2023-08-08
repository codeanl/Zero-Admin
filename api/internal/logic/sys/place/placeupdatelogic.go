package place

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceUpdateLogic) PlaceUpdate(req *types.UpdatePlaceReq) (resp *types.UpdatePlaceResp, err error) {
	_, err = l.svcCtx.Sys.PlaceUpdate(l.ctx, &sysclient.PlaceUpdateReq{
		Id:        req.Id,
		Name:      req.Name,
		Place:     req.Place,
		Status:    req.Status,
		Pic:       req.Pic,
		Phone:     req.Phone,
		Principal: req.Principal,
		UpdateBy:  l.ctx.Value("username").(string),
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdatePlaceResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
