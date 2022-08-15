package trace

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// Trace log request time
func Trace(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	t := time.Now()
	defer func() {
		log.Printf("trace: %s, take %s", info.FullMethod, time.Since(t))
	}()
	return handler(ctx, req)
}
