package interceptor

import (
	"proto-test/interceptor/deadline"
	"proto-test/interceptor/trace"

	"proto-test/interceptor/ratelimit"

	"google.golang.org/grpc"
	"google.golang.org/grpc/tap"
)

func GetUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		trace.Trace,
		deadline.Deadline,
	}
}

func GetTapHandler() tap.ServerInHandle {
	return ratelimit.RateLimit
}
