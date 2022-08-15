package deadline

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// Deadline is a unary server interceptor that adds a deadline to the context.
func Deadline(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx, cancelF := context.WithTimeout(ctx, time.Second*10)
	defer cancelF()
	return handler(ctx, req)
}
