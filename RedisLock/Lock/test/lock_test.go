package test

import (
	"common-toolkits-v1/RedisLock/Lock"
	"context"
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	redisMutex, err := Lock.YamlConfigNewRedisMutex("./redis_lock.yaml")
	if err != nil {
		return
	}
	ctx := context.Background()
	lock, err := redisMutex.GetLock(ctx, "2418")
	if err != nil {
		return
	}
	isLock, err := lock.Lock(ctx)
	if err != nil {
		return
	}
	fmt.Println(isLock)
}
