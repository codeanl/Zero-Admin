package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/oms/model"
	"SimplePick-Mall-Server/service/oms/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn),
	}
}