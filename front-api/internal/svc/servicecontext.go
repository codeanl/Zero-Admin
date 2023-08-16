package svc

import (
	"SimplePick-Mall-Server/front-api/internal/config"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sms/rpc/smsclient"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Ums    umsclient.Ums
	Sms    smsclient.Sms
	Pms    pmsclient.Pms
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Ums:    umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Sms:    smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Pms:    pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
	}
}
