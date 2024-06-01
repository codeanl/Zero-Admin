package subject

import (
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectListLogic) SubjectList(req *types.ListSubjectReq) (*types.ListSubjectResp, error) {
	resp, err := l.svcCtx.Sms.SubjectList(l.ctx, &smsclient.SubjectListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListSubjectResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []types.ListSubjectData
	for _, item := range resp.List {
		listUserData := types.ListSubjectData{
			ID:     item.ID,
			Name:   item.Name,
			Pic:    item.Pic,
			Status: item.Status,
			Sort:   item.Sort,
		}
		list = append(list, listUserData)
	}

	return &types.ListSubjectResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
