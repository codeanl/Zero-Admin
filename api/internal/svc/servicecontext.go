package svc

import (
	"SimplePick-Mall-Server/api/internal/config"
	"SimplePick-Mall-Server/api/middleware"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Sys          sysclient.Sys
	Redis        *redis.Redis
	AddSystemLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	newSys := sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))
	return &ServiceContext{
		Config:       c,
		Sys:          newSys,
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
