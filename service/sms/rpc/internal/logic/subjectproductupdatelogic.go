package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductUpdateLogic {
	return &SubjectProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题商品
func (l *SubjectProductUpdateLogic) SubjectProductUpdate(in *sms.SubjectProductUpdateReq) (*sms.SubjectProductUpdateResp, error) {
	err := l.svcCtx.SubjectProductModel.UpdateSubjectProduct(in.ID, &model.SubjectProduct{
		SubjectID: in.SubjectID,
		ProductID: in.ProductID,
		Status:    in.Status,
		Sort:      in.Sort,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}

	return &sms.SubjectProductUpdateResp{}, nil
}
