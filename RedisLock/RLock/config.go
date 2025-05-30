package RLock

import (
	"fmt"
	"math/rand"
	"time"
)

// 默认数值
const (
	defaultRetryTimes = 3
	defaultRetryDelay = 100 * time.Millisecond
	defaultLockExpiry = 10 * time.Second
	defaultWKeyPrefix = "w-lock:"
	defaultRKeyPrefix = "r-lock:"
)

type RWConfig struct {
	retryTimes int           // retry count - 锁的重试次数
	retryDelay time.Duration // retry time - 锁的重试间隔时间
	wKeyPrefix string        // rKeyPrefix - 读锁key的标识
	rKeyPrefix string        // wKeyPrefix - 写锁key的标识
}

// SingleRWLockConfig SingleLock锁的配置
type SingleRWLockConfig struct {
	RWConfig
	expiry time.Duration // lock time - 锁的超时时间
}

// RedRWLockConfig RedLock锁的配置
type RedRWLockConfig struct {
	RWConfig
	expiry time.Duration // lock time - 锁的超时时间
}

// 生成唯一标识
func generateUniqueID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
}
