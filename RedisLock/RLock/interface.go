package RLock

import (
	"context"
	"time"
)

type Interface interface {
	RLock(ctx context.Context) (bool, error)
	UnRLock(ctx context.Context) (bool, error)
	Lock(ctx context.Context) (bool, error)
	UnLock(ctx context.Context) (bool, error)
	RenewRLock(ctx context.Context) (bool, error)
	RenewLock(ctx context.Context) (bool, error)
	RLockTTL(ctx context.Context) (time.Duration, error)
	LockTTL(ctx context.Context) (time.Duration, error)
}
