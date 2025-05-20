package Lock

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	// ErrFailedToAcquireLock 获取锁失败
	ErrFailedToAcquireLock = errors.New("failed to acquire lock")
	// ErrInvalidLock 无效的锁
	ErrInvalidLock = errors.New("invalid lock")
)

// BasicLock 基础锁
type BasicLock struct {
	client redis.UniversalClient // Redis client - Redis服务
	key    string                // lock key - 锁的Key
	value  string                // lock value - 锁的Value
	expiry time.Duration         // lock time - 锁的超时时间
}

// Lock 获取锁
func (l *BasicLock) lock(ctx context.Context) (bool, error) {
	// 使用SETNX原子操作获取锁
	set, err := l.client.SetNX(ctx, l.key, l.value, l.expiry).Result()
	if err != nil {
		return false, fmt.Errorf("redis operation failed: %w", err)
	}
	return set, nil
}

// UnLock 释放锁
func (l *BasicLock) unLock(ctx context.Context) (bool, error) {

	// 使用Lua脚本保证原子性释放锁
	luaScript := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
	`

	res, err := l.client.Eval(ctx, luaScript, []string{l.key}, l.value).Result()
	if err != nil {
		return false, fmt.Errorf("redis operation failed: %w", err)
	}

	return res.(int64) == 1, nil
}

// Renewal 续期锁
func (l *BasicLock) renewal(ctx context.Context) (bool, error) {

	// 使用Lua脚本检查并续期锁
	luaScript := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("PEXPIRE", KEYS[1], ARGV[2])
	else
		return 0
	end
	`

	// 将过期时间转换为毫秒
	expiryMs := int64(l.expiry / time.Millisecond)
	res, err := l.client.Eval(ctx, luaScript, []string{l.key}, l.value, expiryMs).Result()
	if err != nil {
		return false, fmt.Errorf("redis operation failed: %w", err)
	}

	return res.(int64) == 1, nil
}

// TTL 获取锁剩余时间
func (l *BasicLock) ttl(ctx context.Context) (time.Duration, error) {
	ttl, err := l.client.TTL(ctx, l.key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis operation failed: %w", err)
	}
	return ttl, nil
}
