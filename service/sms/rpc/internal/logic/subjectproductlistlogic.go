package logic

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductListLogic {
	return &SubjectProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 专题列表商品
func (l *SubjectProductListLogic) SubjectProductList(in *sms.SubjectProductListReq) (*sms.SubjectProductListResp, error) {
	all, total, err := l.svcCtx.SubjectProductModel.GetSubjectProductList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.SubjectProductListData
	for _, role := range all {
		list = append(list, &sms.SubjectProductListData{
			ID:        int64(role.ID),
			SubjectID: role.SubjectID,
			ProductID: role.ProductID,
			Status:    role.Status,
			Sort:      role.Sort,
		})
	}
	return &sms.SubjectProductListResp{
		Total: total,
		List:  list,
	}, nil
}
