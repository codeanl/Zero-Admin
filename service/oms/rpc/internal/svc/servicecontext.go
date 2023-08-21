package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/oms/model"
	"SimplePick-Mall-Server/service/oms/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	OrderModel        model.OrderModel
	OrderSkuModel     model.OrderSkuModel
	ReturnReasonModel model.ReturnReasonModel
	ReturnApplyModel  model.ReturnApplyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		OrderModel:        model.NewOrderModel(conn),
		OrderSkuModel:     model.NewOrderSkuModel(conn),
		ReturnReasonModel: model.NewReturnReasonModel(conn),
		ReturnApplyModel:  model.NewReturnApplyModel(conn),
	}
}
