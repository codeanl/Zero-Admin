package product

import (
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"

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
		Id:                  resp.ProductInfo.Id,
		CategoryID:          resp.ProductInfo.CategoryID,
		Name:                resp.ProductInfo.Name,
		Pic:                 resp.ProductInfo.Pic,
		ProductSn:           resp.ProductInfo.ProductSn,
		Desc:                resp.ProductInfo.Desc,
		OriginalPrice:       resp.ProductInfo.OriginalPrice,
		Unit:                resp.ProductInfo.Unit,
		Price:               resp.ProductInfo.Price,
		AttributeCategoryID: resp.ProductInfo.AttributeCategoryID,
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

	var AttributeType1 []types.AttributeList
	var AttributeType2 []types.AttributeList
	for _, i := range resp.Attribute {
		if i.Type == "1" {
			var values []types.ValuesList
			for _, j := range i.Values {
				values = append(values, types.ValuesList{
					Value: j.Value,
				})
			}
			AttributeType1 = append(AttributeType1, types.AttributeList{
				ID:                  i.ID,
				AttributeCategoryID: i.AttributeCategoryID,
				Name:                i.Name,
				Type:                i.Type,
				Values:              values,
			})
		}
		if i.Type == "2" {
			var values []types.ValuesList
			for _, j := range i.Values {
				values = append(values, types.ValuesList{
					Value: j.Value,
				})
			}
			AttributeType2 = append(AttributeType2, types.AttributeList{
				ID:                  i.ID,
				AttributeCategoryID: i.AttributeCategoryID,
				Name:                i.Name,
				Type:                i.Type,
				Values:              values,
			})
		}

	}
	Attribute := types.Attribute{
		AttributeType1: AttributeType1,
		AttributeType2: AttributeType2,
	}
	data := types.InfoData{
		ProductInfo:     productInfo,
		SkuList:         SkuList,
		ImgUrl:          resp.ImgUrl,
		IntroduceImgUrl: resp.IntroduceImgUrl,
		Attribute:       Attribute,
	}
	return &types.ProductInfoResp{
		Code:    200,
		Data:    data,
		Message: "success",
	}, nil
}
