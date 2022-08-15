package interceptor

import (
	"context"
	"log"
	"proto-test/errwraper"
	"proto-test/pkg/grpchelper"
	"sync"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

var sm = sync.Map{}

func deadline(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
	return handler(ctx, req)
}
