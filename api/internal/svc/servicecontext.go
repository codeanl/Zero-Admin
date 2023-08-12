package svc

import (
	"SimplePick-Mall-Server/api/internal/config"
	"SimplePick-Mall-Server/api/middleware"
	"SimplePick-Mall-Server/service/pms/rpc/pmsclient"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"SimplePick-Mall-Server/service/ums/rpc/umsclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Sys          sysclient.Sys
	Pms          pmsclient.Pms
	Ums          umsclient.Ums
	Redis        *redis.Redis
	AddSystemLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	newSys := sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))
	newPms := pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc))
	newUms := umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc))
	return &ServiceContext{
		Config:       c,
		Sys:          newSys,
		Pms:          newPms,
		Ums:          newUms,
		Redis:        newRedis,
		AddSystemLog: middleware.NewAddLogMiddleware(newSys).Handle,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
