// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"SimplePick-Mall-Server/service/sms/rpc/internal/logic"
	"SimplePick-Mall-Server/service/sms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/sms/rpc/sms"
)

type SmsServer struct {
	svcCtx *svc.ServiceContext
	sms.UnimplementedSmsServer
}

func NewSmsServer(svcCtx *svc.ServiceContext) *SmsServer {
	return &SmsServer{
		svcCtx: svcCtx,
	}
}

// 添加广告
func (s *SmsServer) HomeAdvertiseAdd(ctx context.Context, in *sms.HomeAdvertiseAddReq) (*sms.HomeAdvertiseAddResp, error) {
	l := logic.NewHomeAdvertiseAddLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseAdd(in)
}

// 广告列表
func (s *SmsServer) HomeAdvertiseList(ctx context.Context, in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	l := logic.NewHomeAdvertiseListLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseList(in)
}

// 更新广告
func (s *SmsServer) HomeAdvertiseUpdate(ctx context.Context, in *sms.HomeAdvertiseUpdateReq) (*sms.HomeAdvertiseUpdateResp, error) {
	l := logic.NewHomeAdvertiseUpdateLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseUpdate(in)
}

// 删除广告
func (s *SmsServer) HomeAdvertiseDelete(ctx context.Context, in *sms.HomeAdvertiseDeleteReq) (*sms.HomeAdvertiseDeleteResp, error) {
	l := logic.NewHomeAdvertiseDeleteLogic(ctx, s.svcCtx)
	return l.HomeAdvertiseDelete(in)
}

// 添加优惠券
func (s *SmsServer) CouponAdd(ctx context.Context, in *sms.CouponAddReq) (*sms.CouponAddResp, error) {
	l := logic.NewCouponAddLogic(ctx, s.svcCtx)
	return l.CouponAdd(in)
}

// 更新优惠券
func (s *SmsServer) CouponUpdate(ctx context.Context, in *sms.CouponUpdateReq) (*sms.CouponUpdateResp, error) {
	l := logic.NewCouponUpdateLogic(ctx, s.svcCtx)
	return l.CouponUpdate(in)
}

// 删除优惠券
func (s *SmsServer) CouponDelete(ctx context.Context, in *sms.CouponDeleteReq) (*sms.CouponDeleteResp, error) {
	l := logic.NewCouponDeleteLogic(ctx, s.svcCtx)
	return l.CouponDelete(in)
}

// 优惠券列表
func (s *SmsServer) CouponList(ctx context.Context, in *sms.CouponListReq) (*sms.CouponListResp, error) {
	l := logic.NewCouponListLogic(ctx, s.svcCtx)
	return l.CouponList(in)
}

// 添加推荐
func (s *SmsServer) HotRecommendAdd(ctx context.Context, in *sms.HotRecommendAddReq) (*sms.HotRecommendAddResp, error) {
	l := logic.NewHotRecommendAddLogic(ctx, s.svcCtx)
	return l.HotRecommendAdd(in)
}

// 推荐列表
func (s *SmsServer) HotRecommendList(ctx context.Context, in *sms.HotRecommendListReq) (*sms.HotRecommendListResp, error) {
	l := logic.NewHotRecommendListLogic(ctx, s.svcCtx)
	return l.HotRecommendList(in)
}

// 更新推荐
func (s *SmsServer) HotRecommendUpdate(ctx context.Context, in *sms.HotRecommendUpdateReq) (*sms.HotRecommendUpdateResp, error) {
	l := logic.NewHotRecommendUpdateLogic(ctx, s.svcCtx)
	return l.HotRecommendUpdate(in)
}

// 删除推荐
func (s *SmsServer) HotRecommendDelete(ctx context.Context, in *sms.HotRecommendDeleteReq) (*sms.HotRecommendDeleteResp, error) {
	l := logic.NewHotRecommendDeleteLogic(ctx, s.svcCtx)
	return l.HotRecommendDelete(in)
}

// 添加专题
func (s *SmsServer) SubjectAdd(ctx context.Context, in *sms.SubjectAddReq) (*sms.SubjectAddResp, error) {
	l := logic.NewSubjectAddLogic(ctx, s.svcCtx)
	return l.SubjectAdd(in)
}

// 更新专题
func (s *SmsServer) SubjectUpdate(ctx context.Context, in *sms.SubjectUpdateReq) (*sms.SubjectUpdateResp, error) {
	l := logic.NewSubjectUpdateLogic(ctx, s.svcCtx)
	return l.SubjectUpdate(in)
}

// 删除专题
func (s *SmsServer) SubjectDelete(ctx context.Context, in *sms.SubjectDeleteAddReq) (*sms.SubjectDeleteResp, error) {
	l := logic.NewSubjectDeleteLogic(ctx, s.svcCtx)
	return l.SubjectDelete(in)
}

// 专题列表
func (s *SmsServer) SubjectList(ctx context.Context, in *sms.SubjectListReq) (*sms.SubjectListResp, error) {
	l := logic.NewSubjectListLogic(ctx, s.svcCtx)
	return l.SubjectList(in)
}

// 添加专题商品
func (s *SmsServer) SubjectProductAdd(ctx context.Context, in *sms.SubjectProductAddReq) (*sms.SubjectProductAddResp, error) {
	l := logic.NewSubjectProductAddLogic(ctx, s.svcCtx)
	return l.SubjectProductAdd(in)
}

// 更新专题商品
func (s *SmsServer) SubjectProductUpdate(ctx context.Context, in *sms.SubjectProductUpdateReq) (*sms.SubjectProductUpdateResp, error) {
	l := logic.NewSubjectProductUpdateLogic(ctx, s.svcCtx)
	return l.SubjectProductUpdate(in)
}

// 删除专题商品
func (s *SmsServer) SubjectProductDelete(ctx context.Context, in *sms.SubjectProductDeleteAddReq) (*sms.SubjectProductDeleteResp, error) {
	l := logic.NewSubjectProductDeleteLogic(ctx, s.svcCtx)
	return l.SubjectProductDelete(in)
}

// 专题列表商品
func (s *SmsServer) SubjectProductList(ctx context.Context, in *sms.SubjectProductListReq) (*sms.SubjectProductListResp, error) {
	l := logic.NewSubjectProductListLogic(ctx, s.svcCtx)
	return l.SubjectProductList(in)
}