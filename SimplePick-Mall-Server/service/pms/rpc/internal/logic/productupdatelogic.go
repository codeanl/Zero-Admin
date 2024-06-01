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
	//更新
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
		MerchantID:          in.MerchantID,
		Sale:                in.Sale,
	})
	if err != nil {
		return nil, err
	}
	//图片
	if len(in.ImgUrl) > 0 {
		l.svcCtx.ProductImgModel.DeleteProductImgBySpuID(in.Id)
		for _, i := range in.ImgUrl {
			l.svcCtx.ProductImgModel.AddProductImg(&model.ProductImg{
				ProductID: in.Id,
				Url:       i,
			})
		}
	}
	//介绍图片
	if len(in.IntroduceImgUrl) > 0 {
		l.svcCtx.ProductIntroduceImgModel.DeleteProductIntroduceImgBySpuID(in.Id)
		for _, i := range in.IntroduceImgUrl {
			l.svcCtx.ProductIntroduceImgModel.AddProductIntroduceImg(&model.ProductIntroduceImg{
				ProductID: in.Id,
				Url:       i,
			})
		}
	}
	//添加属性
	spu, _ := l.svcCtx.ProductModel.GetProductById(in.Id)
	var AttributeValueType1 []*pms.AttributeValueList
	var AttributeValueType2 []*pms.AttributeValueList
	//填充type1/2数据
	for _, i := range in.AttributeValueList {
		attr, _ := l.svcCtx.AttributeModel.GetAttributeByID(i.AttributeID)
		if attr.Type == "1" {
			AttributeValueType1 = append(AttributeValueType1, &pms.AttributeValueList{
				AttributeID: i.AttributeID,
				Value:       i.Value,
			})
		}
		if attr.Type == "2" {
			AttributeValueType2 = append(AttributeValueType2, &pms.AttributeValueList{
				AttributeID: i.AttributeID,
				Value:       i.Value,
			})
		}
	}
	//type1添加
	if len(AttributeValueType1) > 0 {
		for _, i := range AttributeValueType1 {
			l.svcCtx.AttributeValueModel.DeleteAttributeValueByProductIDAndAttrID(in.Id, i.AttributeID)
			for _, j := range i.Value {
				l.svcCtx.AttributeValueModel.AddAttributeValue(&model.AttributeValue{
					ProductID:   int64(spu.ID),
					AttributeID: i.AttributeID,
					Value:       j,
				})
			}
		}
	}
	//type2添加
	if len(AttributeValueType1) > 0 {
		for _, i := range AttributeValueType2 {
			l.svcCtx.AttributeValueModel.DeleteAttributeValueByProductIDAndAttrID(in.Id, i.AttributeID)
			for _, j := range i.Value {
				l.svcCtx.AttributeValueModel.AddAttributeValue(&model.AttributeValue{
					ProductID:   int64(spu.ID),
					AttributeID: i.AttributeID,
					Value:       j,
				})
			}
		}
	}
	if len(AttributeValueType2) > 0 {
		//添加sku
		var result [][]string
		temp := make([]string, len(AttributeValueType2))
		for _, value := range AttributeValueType2[0].Value {
			temp[0] = value
			generateCombinations(AttributeValueType2, 1, temp, &result)
		}
		//
		skuList, _, _ := l.svcCtx.SkuModel.GetSkuList(&pms.SkuListReq{ProductID: in.Id})
		var nowSku []string
		for _, i := range skuList {
			nowSku = append(nowSku, i.Tag)
		}
		var reqSku []string
		for _, values := range result {
			var data []string
			for _, i := range values {
				attrValue, _ := l.svcCtx.AttributeValueModel.GetAttributeValueBySpuIdAndValue(int64(spu.ID), i)
				info1, _ := l.svcCtx.AttributeModel.GetAttributeByID(attrValue.AttributeID)
				nn := fmt.Sprintf(`{"%s": "%s"}`, info1.Name, attrValue.Value)
				data = append(data, nn)
			}
			tag := strings.Join(data, ", ")
			reqSku = append(reqSku, tag)
		}
		//过滤新添加和需要删除的
		a, b := filterArrays(reqSku, nowSku)
		for _, i := range a {
			l.svcCtx.SkuModel.AddSku(&model.ProductSku{
				ProductID:   int64(spu.ID),
				Name:        spu.Name,
				Pic:         spu.Pic,
				SkuSn:       spu.ProductSn,
				Description: spu.Desc,
				Price:       spu.Price,
				Stock:       100,
				Sale:        0,
				Tag:         i,
			})
		}
		for _, i := range b {
			l.svcCtx.SkuModel.DeleteSkuByTag(i)
		}
	}

	return &pms.ProductUpdateResp{}, nil
}

func filterArrays(reqSku, nowSku []string) ([]string, []string) {
	reqSkuMap := make(map[string]bool)
	nowSkuMap := make(map[string]bool)

	// 将 aa 中的元素添加到 aaMap 中
	for _, num := range reqSku {
		reqSkuMap[num] = true
	}
	// 将 bb 中的元素添加到 bbMap 中
	for _, num := range nowSku {
		nowSkuMap[num] = true
	}
	// 过滤 aa 中存在但 bb 中不存在的元素
	aaOnly := make([]string, 0)
	for _, num := range reqSku {
		if _, exists := nowSkuMap[num]; !exists {
			aaOnly = append(aaOnly, num)
		}
	}
	// 过滤 bb 中存在但 aa 中不存在的元素
	bbOnly := make([]string, 0)
	for _, num := range nowSku {
		if _, exists := reqSkuMap[num]; !exists {
			bbOnly = append(bbOnly, num)
		}
	}
	return aaOnly, bbOnly
}
