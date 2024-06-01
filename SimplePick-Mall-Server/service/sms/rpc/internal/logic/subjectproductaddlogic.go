package logic

import (
	"SimplePick-Mall-Server/service/sms/model"
	"context"
	"errors"

	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductAddLogic {
	return &SubjectProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题商品
func (l *SubjectProductAddLogic) SubjectProductAdd(in *sms.SubjectProductAddReq) (*sms.SubjectProductAddResp, error) {
	for _, i := range in.ProductID {
		role := &model.SubjectProduct{
			SubjectID: in.SubjectID,
			ProductID: i,
			Status:    in.Status,
			Sort:      in.Sort,
		}
		err := l.svcCtx.SubjectProductModel.AddSubjectProduct(role)
		if err != nil {
			return nil, errors.New("添加失败")
		}
	}
	return &sms.SubjectProductAddResp{}, nil
}
