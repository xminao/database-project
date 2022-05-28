package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// UniversityRpc服务依赖
	//UniversityRpcConf zrpc.RpcClientConf
	//做yaml映射
	Database struct {
		Type     string
		User     string
		Password string
		Host     string
		Port     int
		Name     string
	}

	CacheRedis cache.CacheConf
}
