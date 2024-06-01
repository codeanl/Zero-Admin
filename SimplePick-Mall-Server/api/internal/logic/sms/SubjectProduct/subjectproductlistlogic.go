package SubjectProduct

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductListLogic {
	return &SubjectProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectProductListLogic) SubjectProductList(req *types.ListSubjectProductReq) (*types.ListSubjectProductResp, error) {
	resp, err := l.svcCtx.Sms.SubjectProductList(l.ctx, &smsclient.SubjectProductListReq{
		Current:   req.Current,
		PageSize:  req.PageSize,
		Status:    req.Status,
		SubjectID: req.SubjectID,
	})
	if err != nil {
		return &types.ListSubjectProductResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []types.ListSubjectProductData
	for _, item := range resp.List {
		product, _ := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{ID: item.ProductID})
		ProductInfo := types.ProductInfoData{
			Id:            product.ProductInfo.Id,
			CategoryID:    product.ProductInfo.CategoryID,
			Name:          product.ProductInfo.Name,
			Pic:           product.ProductInfo.Pic,
			ProductSn:     product.ProductInfo.ProductSn,
			Desc:          product.ProductInfo.Desc,
			OriginalPrice: product.ProductInfo.OriginalPrice,
			Unit:          product.ProductInfo.Unit,
			Price:         product.ProductInfo.Price,
		}
		listUserData := types.ListSubjectProductData{
			Id:          item.ID,
			SubjectID:   item.SubjectID,
			ProductID:   item.ProductID,
			Status:      item.Status,
			Sort:        item.Sort,
			ProductInfo: ProductInfo,
		}
		list = append(list, listUserData)
	}

	return &types.ListSubjectProductResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
