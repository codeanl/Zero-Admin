package index

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoLogic {
	return &ProductInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductInfoLogic) ProductInfo(req *types.ProductInfoReq) (*types.ProductInfoResp, error) {
	resp, err := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}
	productInfo := types.ListProductData{
		Id:            resp.ProductInfo.Id,
		CategoryID:    resp.ProductInfo.CategoryID,
		Name:          resp.ProductInfo.Name,
		Pic:           resp.ProductInfo.Pic,
		ProductSn:     resp.ProductInfo.ProductSn,
		Desc:          resp.ProductInfo.Desc,
		OriginalPrice: resp.ProductInfo.OriginalPrice,
		Unit:          resp.ProductInfo.Unit,
		Price:         resp.ProductInfo.Price,
	}
	var SkuList []types.SkuList
	for _, i := range resp.SkuList {
		SkuList = append(SkuList, types.SkuList{
			ID:          i.ID,
			ProductID:   i.ProductID,
			Name:        i.Name,
			Pic:         i.Pic,
			SkuSn:       i.SkuSn,
			Description: i.Description,
			Price:       i.Price,
			Stock:       i.Stock,
			Tag:         i.Tag,
		})
	}
	var SizeList []types.SizeList
	for _, i := range resp.SizeList {
		var sizeValue []types.SizeValue
		for _, j := range i.SizeValue {
			sizeValue = append(sizeValue, types.SizeValue{
				ID:     j.ID,
				SizeID: j.SizeID,
				Name:   j.Value,
			})
		}
		SizeList = append(SizeList, types.SizeList{
			ID:        i.ID,
			Name:      i.Name,
			ProductID: i.ProductID,
			SizeValue: sizeValue,
		})
	}

	data := types.InfoData{
		ProductInfo:    productInfo,
		SkuList:        SkuList,
		ImgUrl:         resp.ImgUrl,
		AttributeValue: resp.AttributeValue,
		SizeList:       SizeList,
	}
	return &types.ProductInfoResp{
		Code:    200,
		Data:    data,
		Message: "success",
	}, nil
}
