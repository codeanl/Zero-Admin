package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加商品
func (l *ProductAddLogic) ProductAdd(in *pms.ProductAddReq) (*pms.ProductAddResp, error) {
	info := model.Product{
		CategoryID:          in.CategoryID,
		Name:                in.Name,
		Pic:                 in.Pic,
		ProductSn:           in.ProductSn,
		Desc:                in.Desc,
		OriginalPrice:       in.OriginalPrice,
		Price:               in.Price,
		Unit:                in.Unit,
		AttributeCategoryID: in.AttributeCategoryID,
	}
	spu, err := l.svcCtx.ProductModel.AddProduct(&info)

	for _, i := range in.ImgUrl {
		l.svcCtx.ProductImgModel.AddProductImg(&model.ProductImg{
			ProductID: int64(spu.ID),
			Url:       i,
		})
	}
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
	// 输出结果
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
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &pms.ProductAddResp{}, nil
}
func generateCombinations(attributeValueList []*pms.AttributeValueList, index int, temp []string, result *[][]string) {
	if index == len(attributeValueList) {
		// 将当前组合添加到结果中
		*result = append(*result, append([]string{}, temp...))
		return
	}
	for _, value := range attributeValueList[index].Value {
		temp[index] = value
		generateCombinations(attributeValueList, index+1, temp, result)
	}
}
