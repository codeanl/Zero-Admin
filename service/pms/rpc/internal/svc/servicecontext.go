package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	ProductModel        model.ProductModel
	CategoryModel       model.CategoryModel
	AttributeModel      model.AttributeModel
	AttributeValueModel model.AttributeValueModel
	SpuAttributeModel   model.SpuAttributeModel
	SkuAttributeModel   model.SkuAttributeModel
	SkuModel            model.SkuModel
	SpuSizeModel        model.SpuSizeModel
	SpuSizeValueModel   model.SpuSizeValueModel
	//AttributeCategoryModel model.AttributeCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		ProductModel:        model.NewProductModel(conn),
		CategoryModel:       model.NewCategoryModel(conn),
		AttributeModel:      model.NewAttributeModel(conn),
		AttributeValueModel: model.NewAttributeValueModel(conn),
		SpuAttributeModel:   model.NewSpuAttributeModel(conn),
		SkuAttributeModel:   model.NewSkuAttributeModel(conn),
		SkuModel:            model.NewSkuModel(conn),
		SpuSizeModel:        model.NewSpuSizeModel(conn),
		SpuSizeValueModel:   model.NewSpuSizeValueModel(conn),
	}
}