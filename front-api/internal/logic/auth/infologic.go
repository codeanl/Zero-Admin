package auth

import (
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (*types.InfoResp, error) {
	//id, _ := l.ctx.Value("memberId").(json.Number).Int64()
	resp, err := l.svcCtx.Ums.MemberInfo(l.ctx, &umsclient.MemberInfoReq{Id: req.ID})
	if err != nil {
		return nil, err
	}
	info := types.Info{
		Id:        resp.Id,
		Username:  resp.Username,
		Password:  resp.Password,
		Nickname:  resp.Nickname,
		Phone:     resp.Phone,
		Status:    resp.Status,
		Avatar:    resp.Avatar,
		Gender:    resp.Gender,
		Email:     resp.Email,
		City:      resp.City,
		Job:       resp.Job,
		Signature: resp.Signature,
	}

	return &types.InfoResp{
		Code:    200,
		Message: "success",
		Data:    info,
	}, nil
}
