package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/pms/model"
	"SimplePick-Mall-Server/service/pms/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//ProductModel           model.ProductModel
	//BrandModel             model.BrandModel
	CategoryModel       model.CategoryModel
	AttributeModel      model.AttributeModel
	AttributeValueModel model.AttributeValueModel
	//AttributeCategoryModel model.AttributeCategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config: c,
		//ProductModel:           model.NewProductModel(conn),
		//BrandModel:             model.NewBrandModel(conn),
		CategoryModel:       model.NewCategoryModel(conn),
		AttributeModel:      model.NewAttributeModel(conn),
		AttributeValueModel: model.NewAttributeValueModel(conn),
		//AttributeCategoryModel: model.NewAttributeCategoryModel(conn),
	}
}
