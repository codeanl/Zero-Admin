package logic

import (
	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
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
	productInfo := &pms.ProductListData{
		Id:            int64(resp1.ID),
		CategoryID:    resp1.CategoryID,
		Name:          resp1.Name,
		Pic:           resp1.Pic,
		ProductSn:     resp1.ProductSn,
		Desc:          resp1.Desc,
		Price:         resp1.Price,
		OriginalPrice: resp1.OriginalPrice,
		Unit:          resp1.Unit,
	}

	//获取商品的所有属性
	var attrValue []string
	attrValuelist, _ := l.svcCtx.SpuAttributeModel.GetAttributeValueByProductID(in.ID)
	for _, i := range attrValuelist {
		info1, _ := l.svcCtx.AttributeValueModel.GetAttributeValueByID(i.AttributeValueID)
		info2, _ := l.svcCtx.AttributeModel.GetAttributeByID(info1.AttributeID)
		nn := fmt.Sprintf(`{"%s": "%s"}`, info2.Name, info1.Name)
		attrValue = append(attrValue, nn)
	}
	//tag := strings.Join(attrValue, ", ")

	//获取商品下的size
	var sizeList []*pms.SizeList
	size, _ := l.svcCtx.SpuSizeModel.GetSizeBySpuID(in.ID)
	for _, j := range size {
		var sizeValue []*pms.SizeValue
		info, _ := l.svcCtx.SpuSizeValueModel.GetSizeValueBySizeID(int64(j.ID))
		for _, f := range info {
			sizeValue = append(sizeValue, &pms.SizeValue{
				ID:     int64(f.ID),
				SizeID: f.SizeID,
				Value:  f.Value,
			})
		}
		sizeList = append(sizeList, &pms.SizeList{
			ID:        int64(j.ID),
			ProductID: j.ProductID,
			Name:      j.Name,
			SizeValue: sizeValue,
		})
	}
	//获取商品下的图片
	var ImgUrl []string
	img, _ := l.svcCtx.ProductImgModel.GetImgtByProducID(in.ID)
	for _, i := range img {
		ImgUrl = append(ImgUrl, i.Url)
	}

	//获取sku
	var skuList []*pms.SkuListData
	sku, _, _ := l.svcCtx.SkuModel.GetSkuList(&pms.SkuListReq{ProductID: in.ID})
	for _, i := range sku {
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
	log.Print(skuList)
	return &pms.ProductInfoResp{
		ProductInfo:    productInfo,
		ImgUrl:         ImgUrl,
		SkuList:        skuList,
		SizeList:       sizeList,
		AttributeValue: attrValue,
	}, nil
}
