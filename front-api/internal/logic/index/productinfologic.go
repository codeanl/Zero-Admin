package index

import (
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"context"
	"encoding/json"

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
		var data []map[string]string
		json.Unmarshal([]byte("["+i.Tag+"]"), &data)
		var space []types.Specs
		for _, item := range data {
			for key, value := range item {
				space = append(space, types.Specs{
					Name:      key,
					ValueName: value,
				})
			}
		}
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
			Specs:       space,
		})
	}
	var SizeList []types.SizeList
	var AttributeList []types.AttributeLists
	for _, i := range resp.Attribute {
		if i.Type == "1" {
			var value string
			value = i.Values[0].Value
			AttributeList = append(AttributeList, types.AttributeLists{
				Name:  i.Name,
				Value: value,
			})
		}
		if i.Type == "2" {
			var values []types.Values
			for _, j := range i.Values {
				values = append(values, types.Values{Name: j.Value})
			}
			SizeList = append(SizeList, types.SizeList{
				ID:     i.ID,
				Name:   i.Name,
				Values: values,
			})
		}
	}
	data := types.InfoData{
		ProductInfo:     productInfo,
		SkuList:         SkuList,
		ImgUrl:          resp.ImgUrl,
		IntroduceImgUrl: resp.IntroduceImgUrl,
		AttributeList:   AttributeList,
		SizeList:        SizeList,
	}
	return &types.ProductInfoResp{
		Code:    200,
		Data:    data,
		Message: "success",
	}, nil
}
