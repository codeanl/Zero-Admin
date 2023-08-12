package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SysRpc zrpc.RpcClientConf
	PmsRpc zrpc.RpcClientConf
	UmsRpc zrpc.RpcClientConf
	SmsRpc zrpc.RpcClientConf
	Auth   struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Address string
		Pass    string
	}
}
