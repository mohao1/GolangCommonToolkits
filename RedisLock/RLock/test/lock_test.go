package test

import (
	"common-toolkits-v1/RedisLock/RLock"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	redisMutex, err := RLock.YamlConfigNewRWRedisMutex("./redis_lock.yaml")
	if err != nil {
		return
	}
	ctx := context.Background()
	lock, err := redisMutex.GetLock(ctx, "2418")
	if err != nil {
		return
	}
	isLock, err := lock.Lock(ctx)
	fmt.Println(isLock)

	rlock, err := redisMutex.GetLock(ctx, "2418")
	if err != nil {
		return
	}

	isOk, err := rlock.RLock(ctx)
	if err != nil {
		return
	}
	fmt.Println(isOk)

	time.Sleep(8 * time.Second)

	ttl, err := lock.LockTTL(ctx)
	if err != nil {
		return
	}

	fmt.Println(ttl)

	renewLock, err := lock.RenewLock(ctx)
	if err != nil {
		return
	}

	fmt.Println(renewLock)

	ttl, err = lock.LockTTL(ctx)
	if err != nil {
		return
	}

	fmt.Println(ttl)

	unLock, err := lock.UnLock(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(unLock)
}
