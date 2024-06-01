package svc

import (
	"SimplePick-Mall-Server/common/db"
	"SimplePick-Mall-Server/service/ums/model"
	"SimplePick-Mall-Server/service/ums/rpc/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	MemberModel         model.MemberModel
	MemberLoginLogModel model.MemberLoginLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := db.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		MemberModel:         model.NewMemberModel(conn),
		MemberLoginLogModel: model.NewMemberLoginLogModel(conn),
	}
}
