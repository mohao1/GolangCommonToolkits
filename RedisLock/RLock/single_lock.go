package RLock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"sync"
	"time"
)

type SingleRWLock struct {
	basicRWLock
	RWConfig
	mu       sync.Mutex // lock  - 锁的操作的锁
	released bool       // is released - 是否已被释放
}

func NewSingleRWLock(client redis.Client, resource string, rwConfig SingleRWLockConfig) *SingleRWLock {
	readLockKey := rwConfig.rKeyPrefix + resource
	writeLockKey := rwConfig.wKeyPrefix + resource
	return &SingleRWLock{
		basicRWLock: basicRWLock{
			client:       client,
			ReadLockKey:  readLockKey,
			WriteLockKey: writeLockKey,
			value:        generateUniqueID(),
			expiry:       rwConfig.expiry,
		},
		RWConfig: rwConfig.RWConfig,
		released: true,
	}
}

// RLock 获取读锁
func (s *SingleRWLock) RLock(ctx context.Context) (bool, error) {

	for i := 0; i < s.retryTimes; i++ {
		ok, err := s.basicRWLock.rLock(ctx)
		if err != nil {
			return false, err
		}

		if ok {
			s.released = false
			return true, nil
		}

		// 重试间隔添加随机抖动，避免惊群效应
		delay := s.retryDelay + time.Duration(rand.Int63n(int64(s.retryDelay/2)))
		time.Sleep(delay)
	}
	return false, nil
}

// UnRLock 释放读锁
func (s *SingleRWLock) UnRLock(ctx context.Context) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.released {
		return false, nil
	}

	ok, err := s.basicRWLock.unRLock(ctx)
	if err != nil {
		return false, err
	}

	if ok {
		s.released = true
	}

	return ok, nil
}

// Lock 获取写锁
func (s *SingleRWLock) Lock(ctx context.Context) (bool, error) {

	for i := 0; i < s.retryTimes; i++ {

		ok, err := s.basicRWLock.lock(ctx)
		if err != nil {
			return false, err
		}

		if ok {
			s.released = false
			return true, nil
		}

		// 重试间隔添加随机抖动，避免惊群效应
		delay := s.retryDelay + time.Duration(rand.Int63n(int64(s.retryDelay/2)))
		time.Sleep(delay)
	}
	return false, nil
}

// UnLock 释放写锁
func (s *SingleRWLock) UnLock(ctx context.Context) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.released {
		return false, nil
	}

	ok, err := s.basicRWLock.unLock(ctx)
	if err != nil {
		return false, err
	}

	if ok {
		s.released = true
	}

	return ok, nil
}

// RenewRLock 更新读锁时间
func (s *SingleRWLock) RenewRLock(ctx context.Context) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.released {
		return false, ErrInvalidLock
	}

	return s.basicRWLock.renewRLock(ctx)
}

// RenewLock 更新写锁时间
func (s *SingleRWLock) RenewLock(ctx context.Context) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.released {
		return false, ErrInvalidLock
	}

	return s.basicRWLock.renewLock(ctx)
}

// RLockTTL 获取读锁剩余时间
func (s *SingleRWLock) RLockTTL(ctx context.Context) (time.Duration, error) {
	return s.basicRWLock.rLockTTL(ctx)
}

// LockTTL 获取写锁剩余时间
func (s *SingleRWLock) LockTTL(ctx context.Context) (time.Duration, error) {
	return s.basicRWLock.lockTTL(ctx)
}
