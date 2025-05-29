package Lock

import (
	"common-toolkits-v1/ConfigureParser/YamlParser"
	"common-toolkits-v1/logx/logs"
	"context"
	"crypto/tls"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// MutexConfig Mutex配置解析结构
type MutexConfig struct {
	RedisClients []RedisClientConfig `yaml:"RedisClients" required:"true"` // RedisClient配置列表
	LockConfig   LockConfig          `yaml:"LockConfig"`                   // Lock配置文件
	Expiry       time.Duration       `yaml:"Expiry"`                       // 锁的超时时间
}

// LockConfig 日志设置信息
type LockConfig struct {
	RetryTimes int           `yaml:"RetryTimes"` // retry count - 锁的重试次数
	RetryDelay time.Duration `yaml:"RetryDelay"` // retry time - 锁的重试间隔时间
	KeyPrefix  string        `yaml:"KeyPrefix"`  // KeyPrefix - key的标识
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

type RedisMutex struct {
	clients     []redis.UniversalClient // 存储Client的列表
	mutexConfig MutexConfig             // 配置文件
}

func NewRedisMutex(config MutexConfig) *RedisMutex {
	clients := make([]redis.UniversalClient, len(config.RedisClients))
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
		clients[k] = client
	}
	return &RedisMutex{
		clients:     clients,
		mutexConfig: config,
	}
}

func (r *RedisMutex) GetLock(ctx context.Context, resource string) (Interface, error) {
	clientsLen := len(r.clients)
	config := Config{
		retryTimes: r.mutexConfig.LockConfig.RetryTimes,
		retryDelay: r.mutexConfig.LockConfig.RetryDelay,
		keyPrefix:  r.mutexConfig.LockConfig.KeyPrefix,
	}
	expiry := r.mutexConfig.Expiry

	switch {
	case clientsLen == 0:
		return nil, errors.New("no redis clients")
	case clientsLen == 1:
		client := r.clients[0]
		lockConfig := SingleLockConfig{
			Config: config,
			expiry: expiry,
		}
		singleLock := NewSingleLock(client, resource, &lockConfig)
		return singleLock, nil
	default:
		lockConfig := RedLockConfig{
			Config: config,
			expiry: expiry,
		}
		redLock := NewRedLock(r.clients, resource, &lockConfig)
		return redLock, nil
	}
}

func YamlConfigNewRedisMutex(yamlPath string) (*RedisMutex, error) {
	yamlParser := YamlParser.NewYamlParser()
	configParser := MutexConfig{}
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

	if configParser.LockConfig.KeyPrefix == "" {
		configParser.LockConfig.KeyPrefix = defaultKeyPrefix
	}

	redisMutex := NewRedisMutex(configParser)
	return redisMutex, nil
}
