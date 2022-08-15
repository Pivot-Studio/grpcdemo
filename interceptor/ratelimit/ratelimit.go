package ratelimit

import (
	"context"
	"log"
	"sync"

	"github.com/Pivot-Studio/grpcdemo/errwraper"
	"github.com/Pivot-Studio/grpcdemo/pkg/grpchelper"

	"golang.org/x/time/rate"
	"google.golang.org/grpc/tap"
)

var sm = sync.Map{}

// RateLimit is a unary server interceptor that adds ratelimit to the context.
func RateLimit(ctx context.Context, info *tap.Info) (context.Context, error) {
	ip := grpchelper.GetIP(ctx)
	var limiter *rate.Limiter
	limiterI, ok := sm.Load(ip)
	if !ok {
		limiter = rate.NewLimiter(1, 1)
		sm.Store(ip, rate.NewLimiter(1, 1))
	} else {
		limiter = limiterI.(*rate.Limiter)
	}
	allow := limiter.Allow()

	if !allow {
		log.Printf("ip %s is too frequent", ip)
		return nil, errwraper.Simple("too frequent", errwraper.ErrTooFrequent)
	}
	return ctx, nil
}
