package RLock

import (
	"common-toolkits-v1/ConfigureParser/YamlParser"
	"common-toolkits-v1/logx/logs"
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// RWMutexConfig Mutex配置解析结构
type RWMutexConfig struct {
	RedisClients []RedisClientConfig `yaml:"RedisClients" required:"true"` // RedisClient配置列表
	LockConfig   LockConfig          `yaml:"LockConfig"`                   // Lock配置文件
	Expiry       time.Duration       `yaml:"Expiry"`                       // 锁的超时时间
}

// LockConfig 日志设置信息
type LockConfig struct {
	RetryTimes int           `yaml:"RetryTimes"` // retry count - 锁的重试次数
	RetryDelay time.Duration `yaml:"RetryDelay"` // retry time - 锁的重试间隔时间
	WKeyPrefix string        `yaml:"WKeyPrefix"` // rKeyPrefix - 读锁key的标识
	RKeyPrefix string        `yaml:"RKeyPrefix"` // wKeyPrefix - 写锁key的标识
}

// RedisClientConfig RedisClient连接配置
type RedisClientConfig struct {
	Addr      string        `yaml:"Addr" required:"true"`      // 地址
	User      string        `yaml:"User" default:""`           // 用户名称
	Pass      string        `yaml:"Pass" default:""`           // 用户密码
	DB        int           `yaml:"DB" default:"0"`            // 选择的DB数据库
	IsTLS     bool          `yaml:"IsTLS" default:"false"`     //是否设置TLS
	TLSConfig *tls.Config   `yaml:"TLSConfig"`                 //TLS配置
	IsLimiter bool          `yaml:"IsLimiter" default:"false"` //是否设置Limiter
	Limiter   redis.Limiter `yaml:"Limiter"`                   //Limiter配置
}

type RWRedisMutex struct {
	clients     []redis.Client // 存储Client的列表
	mutexConfig RWMutexConfig  // 配置文件
}

func NewRWRedisMutex(config RWMutexConfig) *RWRedisMutex {
	clients := make([]redis.Client, len(config.RedisClients))
	for k, clientConfig := range config.RedisClients {
		option := &redis.Options{
			Addr:     clientConfig.Addr,
			Password: clientConfig.Pass,
			Username: clientConfig.User,
			DB:       clientConfig.DB,
		}
		if clientConfig.IsTLS {
			option.TLSConfig = clientConfig.TLSConfig
		}
		if clientConfig.IsLimiter {
			option.Limiter = clientConfig.Limiter
		}
		client := redis.NewClient(option)
		clients[k] = *client
	}
	return &RWRedisMutex{
		clients:     clients,
		mutexConfig: config,
	}
}

func (r *RWRedisMutex) GetLock(ctx context.Context, resource string) (Interface, error) {
	clientsLen := len(r.clients)
	config := RWConfig{
		retryTimes: r.mutexConfig.LockConfig.RetryTimes,
		retryDelay: r.mutexConfig.LockConfig.RetryDelay,
		wKeyPrefix: r.mutexConfig.LockConfig.WKeyPrefix,
		rKeyPrefix: r.mutexConfig.LockConfig.RKeyPrefix,
	}
	expiry := r.mutexConfig.Expiry

	switch {
	case clientsLen == 0:
		return nil, errors.New("no redis clients")
	case clientsLen == 1:
		client := r.clients[0]
		lockConfig := SingleRWLockConfig{
			RWConfig: config,
			expiry:   expiry,
		}
		singleLock := NewSingleRWLock(client, resource, lockConfig)
		return singleLock, nil
	default:
		lockConfig := RedRWLockConfig{
			RWConfig: config,
			expiry:   expiry,
		}
		redLock := NewRedRWLock(r.clients, resource, lockConfig)
		return redLock, nil
	}
}

func YamlConfigNewRWRedisMutex(yamlPath string) (*RWRedisMutex, error) {
	yamlParser := YamlParser.NewYamlParser()
	configParser := RWMutexConfig{}
	err := yamlParser.ConfigureParser(yamlPath, &configParser)
	if err != nil {
		logs.Errorf("yamlParser err %v", err)
		return nil, err
	}

	// 初始化的数据设置
	if configParser.Expiry == 0 {
		configParser.Expiry = defaultLockExpiry
	}

	if configParser.LockConfig.RetryTimes == 0 {
		configParser.LockConfig.RetryTimes = defaultRetryTimes
	}

	if configParser.LockConfig.RetryDelay == 0 {
		configParser.LockConfig.RetryDelay = defaultRetryDelay
	}

	if configParser.LockConfig.RKeyPrefix == "" {
		configParser.LockConfig.RKeyPrefix = defaultRKeyPrefix
	}

	if configParser.LockConfig.WKeyPrefix == "" {
		configParser.LockConfig.WKeyPrefix = defaultWKeyPrefix
	}

	redisMutex := NewRWRedisMutex(configParser)
	return redisMutex, nil
}
