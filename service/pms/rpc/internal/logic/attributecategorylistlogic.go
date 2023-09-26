package logic

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeCategoryListLogic {
	return &AttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 属性分类列表
func (l *AttributeCategoryListLogic) AttributeCategoryList(in *pms.AttributeCategoryListReq) (*pms.AttributeCategoryListResp, error) {
	all, total, err := l.svcCtx.AttributeCategoryModel.GetAttributeCategoryList()
	if err != nil {
		return nil, err
	}
	var list []*pms.AttributeCategoryListData
	for _, role := range all {
		list = append(list, &pms.AttributeCategoryListData{
			ID:         int64(role.ID),
			Name:       role.Name,
			ParentID:   role.ParentId,
			MerchantID: role.MerchantID,
		})
	}
	return &pms.AttributeCategoryListResp{
		Total: total,
		List:  list,
	}, nil
}
