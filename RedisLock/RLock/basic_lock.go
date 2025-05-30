package RLock

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var (
	// ErrLockNotFound 锁不存在
	ErrLockNotFound = fmt.Errorf("锁不存在")
	// ErrInvalidLock 无效的锁
	ErrInvalidLock = errors.New("invalid lock")
)

// basicRWLock Redis 读写锁结构
type basicRWLock struct {
	client       redis.Client  // Redis服务
	ReadLockKey  string        // rLock key - 读锁的Key
	WriteLockKey string        // wLock key - 写锁的Key
	value        string        // lock value - 锁的Value
	expiry       time.Duration // lock time - 锁的超时时间
}

// RLock 获取读锁
func (rwl *basicRWLock) rLock(ctx context.Context) (bool, error) {
	// 检查是否有写锁
	hasWriteLock, err := rwl.client.Exists(ctx, rwl.WriteLockKey).Result()
	if err != nil {
		return false, err
	}

	if hasWriteLock > 0 {
		return false, nil
	}

	// 执行 Lua 脚本获取读锁
	result, err := rwl.client.Eval(
		ctx,
		acquireReadLockScript,
		[]string{rwl.ReadLockKey, rwl.WriteLockKey},
		rwl.value,
		strconv.Itoa(int(rwl.expiry.Milliseconds())),
	).Int64()

	if err != nil {
		return false, err
	}

	if result > 0 {
		return true, nil
	}
	return false, nil
}

// UnRLock 释放读锁
func (rwl *basicRWLock) unRLock(ctx context.Context) (bool, error) {
	result, err := rwl.client.Eval(
		ctx,
		releaseReadLockScript,
		[]string{rwl.ReadLockKey},
		rwl.value,
	).Int64()

	if err != nil {
		return false, err
	}

	return result > 0, nil
}

// Lock 获取写锁
func (rwl *basicRWLock) lock(ctx context.Context) (bool, error) {

	// 检查是否有读锁或写锁
	hasLocks, err := rwl.client.Exists(ctx, rwl.ReadLockKey, rwl.WriteLockKey).Result()
	if err != nil {
		return false, err
	}

	if hasLocks > 0 {
		return false, nil
	}

	// 执行 Lua 脚本获取写锁
	result, err := rwl.client.Eval(
		ctx,
		acquireWriteLockScript,
		[]string{rwl.WriteLockKey},
		rwl.value,
		strconv.Itoa(int(rwl.expiry.Milliseconds())),
	).Int64()

	if err != nil {
		return false, err
	}

	if result > 0 {
		return true, nil
	}

	return false, nil
}

// UnLock 释放写锁
func (rwl *basicRWLock) unLock(ctx context.Context) (bool, error) {
	result, err := rwl.client.Eval(
		ctx,
		releaseWriteLockScript,
		[]string{rwl.WriteLockKey},
		rwl.value,
	).Int64()

	if err != nil {
		return false, err
	}

	return result > 0, nil
}

// RenewRLock 续期读锁
func (rwl *basicRWLock) renewRLock(ctx context.Context) (bool, error) {
	result, err := rwl.client.Eval(
		ctx,
		renewReadLockScript,
		[]string{rwl.ReadLockKey},
		rwl.value,
		strconv.Itoa(int(rwl.expiry.Milliseconds())),
	).Int64()

	if err != nil {
		return false, err
	}

	return result > 0, nil
}

// RenewLock 续期写锁
func (rwl *basicRWLock) renewLock(ctx context.Context) (bool, error) {
	result, err := rwl.client.Eval(
		ctx,
		renewWriteLockScript,
		[]string{rwl.WriteLockKey},
		rwl.value,
		strconv.Itoa(int(rwl.expiry.Milliseconds())),
	).Int64()

	if err != nil {
		return false, err
	}

	return result > 0, nil
}

// RLockTTL 获取读锁剩余时间
func (rwl *basicRWLock) rLockTTL(ctx context.Context) (time.Duration, error) {
	ttl, err := rwl.client.Eval(
		ctx,
		getReadLockTTLScript,
		[]string{rwl.ReadLockKey},
		rwl.value,
	).Int64()

	if err != nil {
		return 0, err
	}

	if ttl < 0 {
		return 0, ErrLockNotFound
	}

	return time.Duration(ttl) * time.Millisecond, nil
}

// LockTTL 获取写锁剩余时间
func (rwl *basicRWLock) lockTTL(ctx context.Context) (time.Duration, error) {
	ttl, err := rwl.client.Eval(
		ctx,
		getWriteLockTTLScript,
		[]string{rwl.WriteLockKey},
		rwl.value,
	).Int64()

	if err != nil {
		return 0, err
	}

	if ttl < 0 {
		return 0, ErrLockNotFound
	}

	return time.Duration(ttl) * time.Millisecond, nil
}
