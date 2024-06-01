package logic

import (
	"context"
	"errors"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyDeleteLogic {
	return &MerchantsApplyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商家申请
func (l *MerchantsApplyDeleteLogic) MerchantsApplyDelete(in *pms.MerchantsApplyDeleteReq) (*pms.MerchantsApplyDeleteResp, error) {
	err := l.svcCtx.MerchantsApplyModel.DeleteMerchantsApplyByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.MerchantsApplyDeleteResp{}, nil
}
