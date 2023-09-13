package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type ProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *ProductUpdateLogic) ProductUpdate(in *pms.ProductUpdateReq) (*pms.ProductUpdateResp, error) {
	err := l.svcCtx.ProductModel.UpdateProduct(in.Id, &model.Product{
		CategoryID:          in.CategoryID,
		Name:                in.Name,
		Pic:                 in.Pic,
		ProductSn:           in.ProductSn,
		Desc:                in.Desc,
		OriginalPrice:       in.OriginalPrice,
		Unit:                in.Unit,
		Price:               in.Price,
		AttributeCategoryID: in.AttributeCategoryID,
	})
	if err != nil {
		return nil, err
	}
	l.svcCtx.ProductImgModel.DeleteProductImgBySpuID(in.Id)
	for _, i := range in.ImgUrl {
		l.svcCtx.ProductImgModel.AddProductImg(&model.ProductImg{
			ProductID: in.Id,
			Url:       i,
		})
	}
	l.svcCtx.ProductIntroduceImgModel.DeleteProductIntroduceImgBySpuID(in.Id)
	for _, i := range in.IntroduceImgUrl {
		l.svcCtx.ProductIntroduceImgModel.AddProductIntroduceImg(&model.ProductIntroduceImg{
			ProductID: in.Id,
			Url:       i,
		})
	}
	//
	spu, _ := l.svcCtx.ProductModel.GetProductById(in.Id)
	var AttributeValueType2 []*pms.AttributeValueList
	for _, i := range in.AttributeValueList {
		//添加属性值表
		for _, j := range i.Value {
			l.svcCtx.AttributeValueModel.AddAttributeValue(&model.AttributeValue{
				ProductID:   int64(spu.ID),
				AttributeID: i.AttributeID,
				Value:       j,
			})
		}
		attr, _ := l.svcCtx.AttributeModel.GetAttributeByID(i.AttributeID)
		if attr.Type == "2" {
			AttributeValueType2 = append(AttributeValueType2, &pms.AttributeValueList{
				AttributeID: i.AttributeID,
				Value:       i.Value,
			})
		}
	}
	var result [][]string
	temp := make([]string, len(AttributeValueType2))
	for _, value := range AttributeValueType2[0].Value {
		temp[0] = value
		generateCombinations(AttributeValueType2, 1, temp, &result)
	}
	l.svcCtx.SkuModel.DeleteSkuBySpuID(in.Id)
	for _, values := range result {
		var data []string
		for _, i := range values {
			attrValue, _ := l.svcCtx.AttributeValueModel.GetAttributeValueBySpuIdAndValue(int64(spu.ID), i)
			info1, _ := l.svcCtx.AttributeModel.GetAttributeByID(attrValue.AttributeID)
			nn := fmt.Sprintf(`{"%s": "%s"}`, info1.Name, attrValue.Value)
			data = append(data, nn)
		}
		tag := strings.Join(data, ", ")
		l.svcCtx.SkuModel.AddSku(&model.Sku{
			ProductID:   int64(spu.ID),
			Name:        spu.Name,
			Pic:         spu.Pic,
			SkuSn:       spu.ProductSn,
			Description: spu.Desc,
			Price:       spu.Price,
			Stock:       0,
			Sale:        0,
			Tag:         tag,
		})
	}
	return &pms.ProductUpdateResp{}, nil
}
