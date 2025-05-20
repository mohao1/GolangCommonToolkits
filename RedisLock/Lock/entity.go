package Lock

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
	defaultKeyPrefix  = "lock:"
)

type Config struct {
	retryTimes int           // retry count - 锁的重试次数
	retryDelay time.Duration // retry time - 锁的重试间隔时间
	keyPrefix  string        // KeyPrefix - key的标识
}

// SingleLockConfig SingleLock锁的配置
type SingleLockConfig struct {
	Config
	expiry time.Duration // lock time - 锁的超时时间
}

// NewDefaultLockConfig 创建默认LockConfig
func NewDefaultLockConfig() *SingleLockConfig {
	return &SingleLockConfig{
		Config: Config{
			retryTimes: defaultRetryTimes,
			retryDelay: defaultRetryDelay,
			keyPrefix:  defaultKeyPrefix,
		},
		expiry: defaultLockExpiry,
	}
}

// RedLockConfig RedLock锁的配置
type RedLockConfig struct {
	expiry     time.Duration // lock time - 锁的超时时间
	retryTimes int           // retry count - 锁的重试次数
	retryDelay time.Duration // retry time - 锁的重试间隔时间
	keyPrefix  string        // KeyPrefix - key的标识
}

// 生成唯一标识
func generateUniqueID() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
}
