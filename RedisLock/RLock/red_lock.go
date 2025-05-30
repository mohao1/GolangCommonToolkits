package RLock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const defaultClockDriftFactor = 0.01 // 时钟漂移因子

type RedRWLock struct {
	basicRWLocks []*basicRWLock
	RWConfig
	expiry        time.Duration  // lock time - 设置锁的超时时间
	quorum        int            // 法定人数
	mu            sync.Mutex     // lock  - 锁的操作的锁
	released      bool           // is released - 是否已被释放
	completeLocks []*basicRWLock // 写锁建立上锁连接列表
}

func NewRedRWLock(clients []redis.Client, resource string, rwConfig RedRWLockConfig) *RedRWLock {
	basicRWLocks := make([]*basicRWLock, len(clients))
	for i, client := range clients {
		readLockKey := rwConfig.rKeyPrefix + resource
		writeLockKey := rwConfig.wKeyPrefix + resource
		basicRWLocks[i] = &basicRWLock{
			client:       client,
			ReadLockKey:  readLockKey,
			WriteLockKey: writeLockKey,
			value:        generateUniqueID(),
			expiry:       rwConfig.expiry,
		}
	}
	return &RedRWLock{
		basicRWLocks:  basicRWLocks,
		RWConfig:      rwConfig.RWConfig,
		expiry:        rwConfig.expiry,
		quorum:        (len(clients) / 2) + 1,
		released:      true,
		completeLocks: nil,
	}
}

func (r *RedRWLock) RLock(ctx context.Context) (bool, error) {
	if r.completeLocks != nil {
		return false, errors.New("red_lock already used")
	}

	startTime := time.Now()
	for i := 0; i < r.retryTimes; i++ {
		// 重置
		basicLocks := make([]*basicRWLock, 0, len(r.basicRWLocks))

		// 尝试在所有Redis节点获取锁
		for _, basicLock := range r.basicRWLocks {
			// 获取Lock的锁
			ok, err := basicLock.rLock(ctx)
			if err != nil {
				continue
			}

			if ok {
				basicLocks = append(basicLocks, basicLock)
			}
		}

		// 计算时钟漂移和网络延迟
		drift := time.Duration(float64(r.expiry.Nanoseconds())*defaultClockDriftFactor) + time.Since(startTime)
		validityTime := r.expiry - drift

		// 检查是否达到法定人数且锁有效时间足够
		if len(basicLocks) >= r.quorum && validityTime > 0 {
			r.released = false
			r.completeLocks = basicLocks
			return true, nil
		}

		// 获取失败，释放已获取的锁
		for _, basicLock := range basicLocks {
			_, _ = basicLock.unRLock(ctx)
		}

		if i < r.retryTimes-1 {
			select {
			case <-time.After(r.retryDelay):
			case <-ctx.Done():
				return false, ctx.Err()
			}
		}
	}
	return false, nil
}

func (r *RedRWLock) UnRLock(ctx context.Context) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return false, nil
	}

	successCount := 0

	for _, basicLock := range r.completeLocks {
		ok, err := basicLock.unRLock(ctx)
		if err != nil {
			continue
		}

		if ok {
			successCount++
		}
	}

	completeLocksLen := len(r.completeLocks)
	locksLen := len(r.basicRWLocks)
	if (completeLocksLen - successCount) < (locksLen / 2) {
		r.released = true
		r.completeLocks = nil
		return true, nil
	}

	return false, nil
}

func (r *RedRWLock) Lock(ctx context.Context) (bool, error) {
	if r.completeLocks != nil {
		return false, errors.New("red_lock already used")
	}
	startTime := time.Now()
	for i := 0; i < r.retryTimes; i++ {
		// 重置
		basicLocks := make([]*basicRWLock, 0, len(r.basicRWLocks))

		// 尝试在所有Redis节点获取锁
		for _, basicLock := range r.basicRWLocks {
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
		drift := time.Duration(float64(r.expiry.Nanoseconds())*defaultClockDriftFactor) + time.Since(startTime)
		validityTime := r.expiry - drift

		// 检查是否达到法定人数且锁有效时间足够
		if len(basicLocks) >= r.quorum && validityTime > 0 {
			r.released = false
			r.completeLocks = basicLocks
			return true, nil
		}

		// 获取失败，释放已获取的锁
		for _, basicLock := range basicLocks {
			_, _ = basicLock.unLock(ctx)
		}

		if i < r.retryTimes-1 {
			select {
			case <-time.After(r.retryDelay):
			case <-ctx.Done():
				return false, ctx.Err()
			}
		}
	}
	return false, nil
}

func (r *RedRWLock) UnLock(ctx context.Context) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return false, nil
	}

	successCount := 0

	for _, basicLock := range r.completeLocks {
		ok, err := basicLock.unLock(ctx)
		if err != nil {
			continue
		}

		if ok {
			successCount++
		}
	}

	completeLocksLen := len(r.completeLocks)
	locksLen := len(r.basicRWLocks)
	if (completeLocksLen - successCount) < (locksLen / 2) {
		r.released = true
		r.completeLocks = nil
		return true, nil
	}

	return false, nil
}

func (r *RedRWLock) RenewRLock(ctx context.Context) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return false, ErrInvalidLock
	}

	successCount := 0

	for _, basicLock := range r.completeLocks {
		ok, err := basicLock.renewRLock(ctx)
		if err != nil {
			continue
		}
		if ok {
			successCount++
		}
	}

	return successCount > r.quorum, nil
}

func (r *RedRWLock) RenewLock(ctx context.Context) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return false, ErrInvalidLock
	}

	successCount := 0

	for _, basicLock := range r.completeLocks {
		ok, err := basicLock.renewLock(ctx)
		if err != nil {
			continue
		}
		if ok {
			successCount++
		}
	}

	return successCount > r.quorum, nil
}

func (r *RedRWLock) RLockTTL(ctx context.Context) (time.Duration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return 0, ErrInvalidLock
	}

	var minTTL time.Duration
	first := true

	for _, basicLock := range r.completeLocks {
		ttl, err := basicLock.rLockTTL(ctx)
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

func (r *RedRWLock) LockTTL(ctx context.Context) (time.Duration, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.released {
		return 0, ErrInvalidLock
	}

	var minTTL time.Duration
	first := true

	for _, basicLock := range r.completeLocks {
		ttl, err := basicLock.lockTTL(ctx)
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
