package index

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

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
		Status:   "1",
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
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
		Message: "查询角色列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
