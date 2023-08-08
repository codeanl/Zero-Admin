package svc

import (
	"SimplePick-Mall-Server/api/internal/config"
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Sys    sysclient.Sys
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,
		Sys:    sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Redis:  newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
