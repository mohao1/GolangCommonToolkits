package Lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

// Lock 分布式锁结构
type Lock struct {
	BasicLock
	Config
}

// NewLockDefault 创建新的锁实例 - 默认的值
func NewLockDefault(client redis.UniversalClient, resource string) *Lock {
	config := NewDefaultLockConfig()
	key := config.keyPrefix + resource
	value := generateUniqueID()
	return &Lock{
		BasicLock: BasicLock{
			client:   client,
			key:      key,
			value:    value,
			released: false,
			expiry:   defaultLockExpiry,
		},
		Config: Config{
			retryTimes: defaultRetryTimes,
			retryDelay: defaultRetryDelay,
			keyPrefix:  defaultKeyPrefix,
		},
	}
}

// NewLock 创建新的锁实例 - 设置Config
func NewLock(client redis.UniversalClient, resource string, config *SingleLockConfig) *Lock {
	key := config.keyPrefix + resource
	value := generateUniqueID()
	return &Lock{
		BasicLock: BasicLock{
			client:   client,
			key:      key,
			value:    value,
			released: false,
			expiry:   config.expiry,
		},
		Config: config.Config,
	}
}

// Lock 获取锁
func (l *Lock) Lock(ctx context.Context) (bool, error) {

	for i := 0; i < l.retryTimes; i++ {

		ok, err := l.BasicLock.lock(ctx)
		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}

		// 重试间隔添加随机抖动，避免惊群效应
		delay := l.retryDelay + time.Duration(rand.Int63n(int64(l.retryDelay/2)))
		time.Sleep(delay)
	}

	return false, ErrFailedToAcquireLock
}

// UnLock 释放锁
func (l *Lock) UnLock(ctx context.Context) (bool, error) {
	return l.BasicLock.unLock(ctx)
}

// Renewal 续期锁
func (l *Lock) Renewal(ctx context.Context) (bool, error) {
	return l.BasicLock.renewal(ctx)
}

// TTL 获取锁剩余时间
func (l *Lock) TTL(ctx context.Context) (time.Duration, error) {
	return l.BasicLock.ttl(ctx)
}
