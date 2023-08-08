package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/sys/model"
	"SimplePick-Mall-Server/service/sys/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserRoleModel model.UserRoleModel
	RoleModel     model.RoleModel
	LoginLogModel model.LoginLogModel
	PlaceModel    model.PlaceModel
	MenuModel     model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn),
		UserRoleModel: model.NewUserRoleModel(conn),
		RoleModel:     model.NewRoleModel(conn),
		LoginLogModel: model.NewLoginLogModel(conn),
		PlaceModel:    model.NewPlaceModel(conn),
		MenuModel:     model.NewMenuModel(conn),
	}
}
