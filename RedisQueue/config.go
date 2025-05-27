package RedisQueue

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
)

// RedisConfig RedisClient配置
type RedisConfig struct {
	Addr      string
	Pass      string
	DB        int
	IsTLS     bool
	TLSConfig *tls.Config
	IsLimiter bool
	Limiter   redis.Limiter
}
