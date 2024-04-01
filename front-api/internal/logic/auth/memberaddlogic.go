package auth

import (
	"SimplePick-Mall-Server/common/errorx"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddLogic) MemberAdd(req *types.MemberLoginReq, ip string) (*types.MemberLoginResp, error) {
	resp, err := l.svcCtx.Ums.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewCodeError(400, "查询用户异常")
	}
	l.svcCtx.Ums.MemberLoginLogAdd(l.ctx, &umsclient.MemberLoginLogAddReq{
		UserID: resp.Member.Id,
		IP:     ip,
	})
	if resp.Member.Status == "0" {
		return nil, errorx.NewCodeError(400, "用户已锁定")
	}
	userInfo := types.Info{
		Id:        resp.Member.Id,
		Username:  resp.Member.Username,
		Password:  resp.Member.Password,
		Nickname:  resp.Member.Nickname,
		Phone:     resp.Member.Phone,
		Status:    resp.Member.Status,
		Avatar:    resp.Member.Avatar,
		Gender:    resp.Member.Gender,
		Email:     resp.Member.Email,
		City:      resp.Member.City,
		Job:       resp.Member.Job,
		Signature: resp.Member.Signature,
		CreatTIme: resp.Member.CreateTime,
	}
	loginData := types.LoginData{
		Token: resp.Token,
		Info:  userInfo,
	}
	return &types.MemberLoginResp{
		Code:    200,
		Message: "success",
		Data:    loginData,
	}, nil
}
