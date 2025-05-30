package Lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"sync"
	"time"
)

// SingleLock 分布式锁结构
type SingleLock struct {
	BasicLock
	Config
	mu       sync.Mutex // lock  - 锁的操作的锁
	released bool       // is released - 是否已被释放
}

// NewSingleLockDefault 创建新的锁实例 - 默认的值
func NewSingleLockDefault(client redis.UniversalClient, resource string) *SingleLock {
	key := defaultKeyPrefix + resource
	value := generateUniqueID()
	return &SingleLock{
		BasicLock: BasicLock{
			client: client,
			key:    key,
			value:  value,
			expiry: defaultLockExpiry,
		},
		Config: Config{
			retryTimes: defaultRetryTimes,
			retryDelay: defaultRetryDelay,
			keyPrefix:  defaultKeyPrefix,
		},
		released: true,
	}
}

// NewSingleLock 创建新的锁实例 - 设置Config
func NewSingleLock(client redis.UniversalClient, resource string, config *SingleLockConfig) *SingleLock {
	key := config.keyPrefix + resource
	value := generateUniqueID()
	return &SingleLock{
		BasicLock: BasicLock{
			client: client,
			key:    key,
			value:  value,
			expiry: config.expiry,
		},
		Config:   config.Config,
		released: true,
	}
}

// Lock 获取锁
func (l *SingleLock) Lock(ctx context.Context) (bool, error) {

	for i := 0; i < l.retryTimes; i++ {

		ok, err := l.BasicLock.lock(ctx)
		if err != nil {
			return false, err
		}

		if ok {
			l.released = false
			return true, nil
		}

		// 重试间隔添加随机抖动，避免惊群效应
		delay := l.retryDelay + time.Duration(rand.Int63n(int64(l.retryDelay/2)))
		time.Sleep(delay)
	}

	return false, nil
}

// UnLock 释放锁
func (l *SingleLock) UnLock(ctx context.Context) (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.released {
		return false, nil
	}

	ok, err := l.BasicLock.unLock(ctx)
	if err != nil {
		return false, err
	}

	if ok {
		l.released = true
	}

	return ok, nil
}

// Renewal 续期锁
func (l *SingleLock) Renewal(ctx context.Context) (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.released {
		return false, ErrInvalidLock
	}

	return l.BasicLock.renewal(ctx)
}

// TTL 获取锁剩余时间
func (l *SingleLock) TTL(ctx context.Context) (time.Duration, error) {
	return l.BasicLock.ttl(ctx)
}
