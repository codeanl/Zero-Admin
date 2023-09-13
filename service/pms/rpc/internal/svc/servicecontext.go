package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/config"
)

type ServiceContext struct {
	Config                   config.Config
	ProductModel             model.ProductModel
	CategoryModel            model.CategoryModel
	AttributeModel           model.AttributeModel
	AttributeValueModel      model.AttributeValueModel
	SkuModel                 model.SkuModel
	ProductImgModel          model.ProductImgModel
	AttributeCategoryModel   model.AttributeCategoryModel
	ProductIntroduceImgModel model.ProductIntroduceImgModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:                   c,
		ProductModel:             model.NewProductModel(conn),
		CategoryModel:            model.NewCategoryModel(conn),
		AttributeModel:           model.NewAttributeModel(conn),
		AttributeValueModel:      model.NewAttributeValueModel(conn),
		SkuModel:                 model.NewSkuModel(conn),
		ProductImgModel:          model.NewProductImgModel(conn),
		AttributeCategoryModel:   model.NewAttributeCategoryModel(conn),
		ProductIntroduceImgModel: model.NewProductIntroduceImgModel(conn),
	}
}
