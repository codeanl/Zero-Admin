package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoLogic {
	return &ProductInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询商品详情
func (l *ProductInfoLogic) ProductInfo(in *pms.ProductInfoReq) (*pms.ProductInfoResp, error) {
	resp1, _ := l.svcCtx.ProductModel.GetProductById(in.ID)

	var attrIDs []int64
	attrValue, _ := l.svcCtx.AttributeValueModel.GetAttributeValueBySPUID(in.ID)
	for _, i := range attrValue {
		attrIDs = append(attrIDs, i.AttributeID)
	}
	filteredNums := FilterDuplicates(attrIDs)
	var Attribute []*pms.Attribute
	for _, i := range filteredNums {
		attr, _ := l.svcCtx.AttributeModel.GetAttributeByID(i)
		li, _ := l.svcCtx.AttributeValueModel.GetAttributeValueBySpuIdAndAttrId(in.ID, i)
		var value []*pms.Values
		for _, j := range li {
			value = append(value, &pms.Values{
				Value: j.Value,
			})
		}
		Attribute = append(Attribute, &pms.Attribute{
			ID:                  int64(attr.ID),
			AttributeCategoryID: attr.AttributeCategoryID,
			Name:                attr.Name,
			Type:                attr.Type,
			Values:              value,
		})
	}
	//获取商品下的图片
	var ImgUrl []string
	img, _ := l.svcCtx.ProductImgModel.GetImgtByProducID(in.ID)
	for _, i := range img {
		ImgUrl = append(ImgUrl, i.Url)
	}
	var IntroduceImgUrl []string
	introduceImgUrl, _ := l.svcCtx.ProductIntroduceImgModel.GetImgtByProducID(in.ID)
	for _, i := range introduceImgUrl {
		IntroduceImgUrl = append(IntroduceImgUrl, i.Url)
	}
	//获取sku
	var skuList []*pms.SkuListData
	var sale int64 = 0
	sku, _, _ := l.svcCtx.SkuModel.GetSkuList(&pms.SkuListReq{ProductID: in.ID})
	for _, i := range sku {
		sale = sale + i.Sale
		skuList = append(skuList, &pms.SkuListData{
			ID:          int64(i.ID),
			ProductID:   i.ProductID,
			Name:        i.Name,
			Pic:         i.Pic,
			SkuSn:       i.SkuSn,
			Description: i.Description,
			Price:       i.Price,
			Stock:       i.Stock,
			Sale:        i.Sale,
			Tag:         i.Tag,
		})
	}
	productInfo := &pms.ProductListData{
		Id:                  int64(resp1.ID),
		CategoryID:          resp1.CategoryID,
		Name:                resp1.Name,
		Pic:                 resp1.Pic,
		ProductSn:           resp1.ProductSn,
		Desc:                resp1.Desc,
		Price:               resp1.Price,
		OriginalPrice:       resp1.OriginalPrice,
		Unit:                resp1.Unit,
		AttributeCategoryID: resp1.AttributeCategoryID,
		MerchantID:          resp1.MerchantID,
		Sale:                sale,
	}
	return &pms.ProductInfoResp{
		ProductInfo:     productInfo,
		ImgUrl:          ImgUrl,
		IntroduceImgUrl: IntroduceImgUrl,
		SkuList:         skuList,
		Attribute:       Attribute,
	}, nil
}
func FilterDuplicates(nums []int64) []int64 {
	uniqueNums := make(map[int64]bool)
	var result []int64
	for _, num := range nums {
		if !uniqueNums[num] {
			uniqueNums[num] = true
			result = append(result, num)
		}
	}
	return result
}
