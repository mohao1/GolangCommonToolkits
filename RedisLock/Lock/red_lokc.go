package Lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const defaultClockDriftFactor = 0.01 // 时钟漂移因子

// RedLock 实现RedLock算法（多节点Redis分布式锁）
type RedLock struct {
	basicLocks []*BasicLock
	Config
	expiry   time.Duration // lock time - 设置锁的超时时间
	quorum   int           // 法定人数
	mu       sync.Mutex    // lock  - 锁的操作的锁
	released bool          // is released - 是否已被释放
}

// NewRedLockDefault 创建新的锁实例 - 默认的值
func NewRedLockDefault(clients []redis.UniversalClient, resource string) *RedLock {

	basicLocks := make([]*BasicLock, len(clients))
	key := defaultKeyPrefix + resource
	value := generateUniqueID()

	for i, client := range clients {
		basicLocks[i] = &BasicLock{
			client: client,
			key:    key,
			value:  value,
			expiry: defaultLockExpiry,
		}
	}

	return &RedLock{
		basicLocks: basicLocks,
		Config: Config{
			retryTimes: defaultRetryTimes,
			retryDelay: defaultRetryDelay,
			keyPrefix:  defaultKeyPrefix,
		},
		expiry:   defaultLockExpiry,
		quorum:   (len(clients) / 2) + 1,
		released: false,
	}
}

// NewRedLock 创建新的锁实例 - 设置Config
func NewRedLock(clients []redis.UniversalClient, resource string, config *RedLockConfig) *RedLock {

	basicLocks := make([]*BasicLock, len(clients))
	key := config.keyPrefix + resource
	value := generateUniqueID()

	for i, client := range clients {
		basicLocks[i] = &BasicLock{
			client: client,
			key:    key,
			value:  value,
			expiry: config.expiry,
		}
	}

	return &RedLock{
		basicLocks: basicLocks,
		Config:     config.Config,
		expiry:     config.expiry,
		quorum:     (len(clients) / 2) + 1,
		released:   false,
	}
}

// Lock 使用RedLock算法获取锁
func (rl *RedLock) Lock(ctx context.Context) (bool, error) {
	startTime := time.Now()

	for i := 0; i < rl.retryTimes; i++ {
		// 重置
		basicLocks := make([]*BasicLock, 0, len(rl.basicLocks))

		// 尝试在所有Redis节点获取锁
		for _, basicLock := range rl.basicLocks {
			// 获取Lock的锁
			ok, err := basicLock.lock(ctx)
			if err != nil {
				continue
			}

			if ok {
				basicLocks = append(basicLocks, basicLock)
			}
		}

		// 计算时钟漂移和网络延迟
		drift := time.Duration(float64(rl.expiry.Nanoseconds())*defaultClockDriftFactor) + time.Since(startTime)
		validityTime := rl.expiry - drift

		// 检查是否达到法定人数且锁有效时间足够
		if len(basicLocks) >= rl.quorum && validityTime > 0 {
			return true, nil
		}

		// 获取失败，释放已获取的锁
		for _, basicLock := range basicLocks {
			_, _ = basicLock.unLock(ctx)
		}

		if i < rl.retryTimes-1 {
			select {
			case <-time.After(rl.retryDelay):
			case <-ctx.Done():
				return false, ctx.Err()
			}
		}

	}

	return false, ErrFailedToAcquireLock
}

// UnLock 释放锁
func (rl *RedLock) UnLock(ctx context.Context) (bool, error) {

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.released {
		return false, nil
	}

	successCount := 0

	for _, basicLock := range rl.basicLocks {
		ok, err := basicLock.unLock(ctx)
		if err != nil {
			continue
		}

		if ok {
			successCount++
		}
	}

	rl.released = true

	return successCount > 0, nil
}

func (rl *RedLock) Renewal(ctx context.Context) (bool, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.released {
		return false, ErrInvalidLock
	}

	successCount := 0

	for _, basicLock := range rl.basicLocks {
		ok, err := basicLock.renewal(ctx)
		if err != nil {
			continue
		}
		if ok {
			successCount++
		}
	}

	return successCount > rl.quorum, nil
}

func (rl *RedLock) TTL(ctx context.Context) (time.Duration, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.released {
		return 0, ErrInvalidLock
	}

	var minTTL time.Duration
	first := true

	for _, basicLock := range rl.basicLocks {
		ttl, err := basicLock.ttl(ctx)
		if err != nil {
			continue
		}

		if ttl < 0 {
			continue
		}

		if first {
			minTTL = ttl
			first = false
		} else if ttl < minTTL {
			minTTL = ttl
		}
	}

	if first { // 没有有效的TTL值
		return 0, nil
	}

	return minTTL, nil
}
