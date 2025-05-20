package Lock

import (
	"context"
	"time"
)

type Interface interface {
	Lock(ctx context.Context) (bool, error)
	UnLock(ctx context.Context) (bool, error)
	Renewal(ctx context.Context) (bool, error)
	TTL(ctx context.Context) (time.Duration, error)
}
