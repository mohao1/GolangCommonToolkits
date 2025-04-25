package RedisQueue

import (
	"context"
)

// RedisQueue 定义 Redis 队列结构体
type RedisQueue struct {
	client    Interface
	queueName string
}

// NewRedisQueueConfig 设置配置文件创建新的队列
func NewRedisQueueConfig(redisConfig RedisConfig, queueName string) *RedisQueue {
	client, err := NewRedisClient(redisConfig)
	if err != nil {
		return nil
	}

	return &RedisQueue{
		client:    client,
		queueName: queueName,
	}
}

// NewRedisQueueByClient 自定义redisClient创建队列
func NewRedisQueueByClient(redisClient Interface, queueName string) *RedisQueue {
	return &RedisQueue{
		client:    redisClient,
		queueName: queueName,
	}
}

// Enqueue 向队列中添加元素
func (q *RedisQueue) Enqueue(ctx context.Context, value string) (int, error) {
	return q.client.LPushCtx(ctx, q.queueName, value)
}

// Dequeue 从队列中取出元素
func (q *RedisQueue) Dequeue(ctx context.Context) (string, error) {
	return q.client.RPopCtx(ctx, q.queueName)
}

// HasData 判断队列中是否存在数据
func (q *RedisQueue) HasData(ctx context.Context) (bool, error) {
	length, err := q.client.LLenCtx(ctx, q.queueName)
	if err != nil {
		return false, err
	}
	return length > 0, nil
}
