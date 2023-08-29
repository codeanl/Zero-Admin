package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttributeListLogic {
	return &AttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 属性列表
func (l *AttributeListLogic) AttributeList(in *pms.AttributeListReq) (*pms.AttributeListResp, error) {
	all, total, err := l.svcCtx.AttributeModel.GetAttributeList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.AttributeListData
	for _, i := range all {
		list = append(list, &pms.AttributeListData{
			Id:                  int64(i.ID),
			AttributeCategoryID: i.AttributeCategoryID,
			Name:                i.Name,
			Type:                i.Type,
			Value:               i.Value,
			Sort:                i.Sort,
		})
	}
	return &pms.AttributeListResp{
		Total: total,
		List:  list,
	}, nil
}
