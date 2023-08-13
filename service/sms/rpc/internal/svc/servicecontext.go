package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/sms/model"
	"SimplePick-Mall-Server/service/sms/rpc/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	HomeAdvertiseModel model.HomeAdvertiseModel
	CouponModel        model.CouponModel
	HotRecommendModel  model.HotRecommendModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		HomeAdvertiseModel: model.NewHomeAdvertiseModel(conn),
		CouponModel:        model.NewCouponModel(conn),
		HotRecommendModel:  model.NewHotRecommendModel(conn),
	}
}
