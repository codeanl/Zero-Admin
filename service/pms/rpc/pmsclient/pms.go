// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package pmsclient

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AttributeAddReq     = pms.AttributeAddReq
	AttributeAddResp    = pms.AttributeAddResp
	AttributeDeleteReq  = pms.AttributeDeleteReq
	AttributeDeleteResp = pms.AttributeDeleteResp
	AttributeListData   = pms.AttributeListData
	AttributeListReq    = pms.AttributeListReq
	AttributeListResp   = pms.AttributeListResp
	AttributeUpdateReq  = pms.AttributeUpdateReq
	AttributeUpdateResp = pms.AttributeUpdateResp
	AttributeValue      = pms.AttributeValue
	CategoryAddReq      = pms.CategoryAddReq
	CategoryAddResp     = pms.CategoryAddResp
	CategoryDeleteReq   = pms.CategoryDeleteReq
	CategoryDeleteResp  = pms.CategoryDeleteResp
	CategoryListData    = pms.CategoryListData
	CategoryListReq     = pms.CategoryListReq
	CategoryListResp    = pms.CategoryListResp
	CategoryUpdateReq   = pms.CategoryUpdateReq
	CategoryUpdateResp  = pms.CategoryUpdateResp
	ProductAddReq       = pms.ProductAddReq
	ProductAddResp      = pms.ProductAddResp
	ProductDeleteReq    = pms.ProductDeleteReq
	ProductDeleteResp   = pms.ProductDeleteResp
	ProductInfoReq      = pms.ProductInfoReq
	ProductInfoResp     = pms.ProductInfoResp
	ProductListData     = pms.ProductListData
	ProductListReq      = pms.ProductListReq
	ProductListResp     = pms.ProductListResp
	ProductUpdateReq    = pms.ProductUpdateReq
	ProductUpdateResp   = pms.ProductUpdateResp
	Size                = pms.Size
	SizeList            = pms.SizeList
	SizeValue           = pms.SizeValue
	SkuAddReq           = pms.SkuAddReq
	SkuAddResp          = pms.SkuAddResp
	SkuDeleteReq        = pms.SkuDeleteReq
	SkuDeleteResp       = pms.SkuDeleteResp
	SkuListData         = pms.SkuListData
	SkuListReq          = pms.SkuListReq
	SkuListResp         = pms.SkuListResp
	SkuUpdateReq        = pms.SkuUpdateReq
	SkuUpdateResp       = pms.SkuUpdateResp

	Pms interface {
		// 添加分类
		CategoryAdd(ctx context.Context, in *CategoryAddReq, opts ...grpc.CallOption) (*CategoryAddResp, error)
		// 分类列表
		CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		// 更新分类
		CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*CategoryUpdateResp, error)
		// 删除分类
		CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error)
		// 添加属性
		AttributeAdd(ctx context.Context, in *AttributeAddReq, opts ...grpc.CallOption) (*AttributeAddResp, error)
		// 属性列表
		AttributeList(ctx context.Context, in *AttributeListReq, opts ...grpc.CallOption) (*AttributeListResp, error)
		// 更新属性
		AttributeUpdate(ctx context.Context, in *AttributeUpdateReq, opts ...grpc.CallOption) (*AttributeUpdateResp, error)
		// 删除属性
		AttributeDelete(ctx context.Context, in *AttributeDeleteReq, opts ...grpc.CallOption) (*AttributeDeleteResp, error)
		// 添加商品
		ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error)
		// 商品列表
		ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error)
		// 更新商品
		ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error)
		// 删除商品
		ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error)
		// 查询商品详情
		ProductInfo(ctx context.Context, in *ProductInfoReq, opts ...grpc.CallOption) (*ProductInfoResp, error)
		// 添加Sku
		SkuAdd(ctx context.Context, in *SkuAddReq, opts ...grpc.CallOption) (*SkuAddResp, error)
		// Sku列表
		SkuList(ctx context.Context, in *SkuListReq, opts ...grpc.CallOption) (*SkuListResp, error)
		// 更新Sku
		SkuUpdate(ctx context.Context, in *SkuUpdateReq, opts ...grpc.CallOption) (*SkuUpdateResp, error)
		// 删除Sku
		SkuDelete(ctx context.Context, in *SkuDeleteReq, opts ...grpc.CallOption) (*SkuDeleteResp, error)
	}

	defaultPms struct {
		cli zrpc.Client
	}
)

func NewPms(cli zrpc.Client) Pms {
	return &defaultPms{
		cli: cli,
	}
}

// 添加分类
func (m *defaultPms) CategoryAdd(ctx context.Context, in *CategoryAddReq, opts ...grpc.CallOption) (*CategoryAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryAdd(ctx, in, opts...)
}

// 分类列表
func (m *defaultPms) CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryList(ctx, in, opts...)
}

// 更新分类
func (m *defaultPms) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*CategoryUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryUpdate(ctx, in, opts...)
}

// 删除分类
func (m *defaultPms) CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryDelete(ctx, in, opts...)
}

// 添加属性
func (m *defaultPms) AttributeAdd(ctx context.Context, in *AttributeAddReq, opts ...grpc.CallOption) (*AttributeAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeAdd(ctx, in, opts...)
}

// 属性列表
func (m *defaultPms) AttributeList(ctx context.Context, in *AttributeListReq, opts ...grpc.CallOption) (*AttributeListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeList(ctx, in, opts...)
}

// 更新属性
func (m *defaultPms) AttributeUpdate(ctx context.Context, in *AttributeUpdateReq, opts ...grpc.CallOption) (*AttributeUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeUpdate(ctx, in, opts...)
}

// 删除属性
func (m *defaultPms) AttributeDelete(ctx context.Context, in *AttributeDeleteReq, opts ...grpc.CallOption) (*AttributeDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeDelete(ctx, in, opts...)
}

// 添加商品
func (m *defaultPms) ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductAdd(ctx, in, opts...)
}

// 商品列表
func (m *defaultPms) ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

// 更新商品
func (m *defaultPms) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductUpdate(ctx, in, opts...)
}

// 删除商品
func (m *defaultPms) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductDelete(ctx, in, opts...)
}

// 查询商品详情
func (m *defaultPms) ProductInfo(ctx context.Context, in *ProductInfoReq, opts ...grpc.CallOption) (*ProductInfoResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductInfo(ctx, in, opts...)
}

// 添加Sku
func (m *defaultPms) SkuAdd(ctx context.Context, in *SkuAddReq, opts ...grpc.CallOption) (*SkuAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuAdd(ctx, in, opts...)
}

// Sku列表
func (m *defaultPms) SkuList(ctx context.Context, in *SkuListReq, opts ...grpc.CallOption) (*SkuListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuList(ctx, in, opts...)
}

// 更新Sku
func (m *defaultPms) SkuUpdate(ctx context.Context, in *SkuUpdateReq, opts ...grpc.CallOption) (*SkuUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuUpdate(ctx, in, opts...)
}

// 删除Sku
func (m *defaultPms) SkuDelete(ctx context.Context, in *SkuDeleteReq, opts ...grpc.CallOption) (*SkuDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuDelete(ctx, in, opts...)
}
