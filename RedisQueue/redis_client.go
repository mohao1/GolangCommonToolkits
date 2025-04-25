package RedisQueue

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-redis/redis/v8"
)

type Interface interface {
	LPushCtx(ctx context.Context, key string, value interface{}) (int, error)
	RPopCtx(ctx context.Context, key string) (string, error)
	LLenCtx(ctx context.Context, key string) (int, error)
}

type RedisConfig struct {
	Addr      string
	Pass      string
	DB        int
	IsTLS     bool
	TLSConfig *tls.Config
	IsLimiter bool
	Limiter   redis.Limiter
}

type RedisClient struct {
	redisClient *redis.Client
	Config      RedisConfig
}

// NewRedisClient 创建Client
func NewRedisClient(config RedisConfig) (*RedisClient, error) {

	if config.Addr == "" {
		return nil, errors.New("redis addr is required")
	}

	opt := &redis.Options{
		Addr:     config.Addr,
		Password: config.Pass,
		DB:       config.DB,
	}

	if config.IsTLS {
		opt.TLSConfig = config.TLSConfig
	}

	if config.IsLimiter {
		opt.Limiter = config.Limiter
	}

	client := redis.NewClient(opt)

	return &RedisClient{
		redisClient: client,
		Config:      config,
	}, nil
}

// LPushCtx 设置数据
func (r *RedisClient) LPushCtx(ctx context.Context, key string, value interface{}) (int, error) {
	cmd := r.redisClient.LPush(ctx, key, value)
	err := cmd.Err()
	if err != nil {
		return 0, err
	}
	val := cmd.Val()
	return int(val), nil
}

// RPopCtx 取出数据
func (r *RedisClient) RPopCtx(ctx context.Context, key string) (string, error) {
	cmd := r.redisClient.RPop(ctx, key)
	err := cmd.Err()
	if err != nil {
		return "", err
	}
	val := cmd.Val()
	return val, nil
}

// LLenCtx 获取长度
func (r *RedisClient) LLenCtx(ctx context.Context, key string) (int, error) {
	cmd := r.redisClient.LLen(ctx, key)
	err := cmd.Err()
	if err != nil {
		return 0, err
	}
	val := cmd.Val()
	return int(val), nil
}
