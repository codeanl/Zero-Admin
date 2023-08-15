// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"SimplePick-Mall-Server/service/ums/rpc/internal/logic"
	"SimplePick-Mall-Server/service/ums/rpc/internal/svc"
	"SimplePick-Mall-Server/service/ums/rpc/ums"
)

type UmsServer struct {
	svcCtx *svc.ServiceContext
	ums.UnimplementedUmsServer
}

func NewUmsServer(svcCtx *svc.ServiceContext) *UmsServer {
	return &UmsServer{
		svcCtx: svcCtx,
	}
}

// 添加会员
func (s *UmsServer) MemberAdd(ctx context.Context, in *ums.MemberAddReq) (*ums.MemberAddResp, error) {
	l := logic.NewMemberAddLogic(ctx, s.svcCtx)
	return l.MemberAdd(in)
}

// 会员登录
func (s *UmsServer) MemberLogin(ctx context.Context, in *ums.MemberLoginReq) (*ums.MemberLoginResp, error) {
	l := logic.NewMemberLoginLogic(ctx, s.svcCtx)
	return l.MemberLogin(in)
}

// 会员列表
func (s *UmsServer) MemberList(ctx context.Context, in *ums.MemberListReq) (*ums.MemberListResp, error) {
	l := logic.NewMemberListLogic(ctx, s.svcCtx)
	return l.MemberList(in)
}

// 会员更新
func (s *UmsServer) MemberUpdate(ctx context.Context, in *ums.MemberUpdateReq) (*ums.MemberUpdateResp, error) {
	l := logic.NewMemberUpdateLogic(ctx, s.svcCtx)
	return l.MemberUpdate(in)
}

// 会员删除
func (s *UmsServer) MemberDelete(ctx context.Context, in *ums.MemberDeleteReq) (*ums.MemberDeleteResp, error) {
	l := logic.NewMemberDeleteLogic(ctx, s.svcCtx)
	return l.MemberDelete(in)
}

// 会员详情
func (s *UmsServer) MemberInfo(ctx context.Context, in *ums.MemberInfoReq) (*ums.MemberInfoResp, error) {
	l := logic.NewMemberInfoLogic(ctx, s.svcCtx)
	return l.MemberInfo(in)
}
