package interceptor

import (
	"github.com/Pivot-Studio/grpcdemo/interceptor/deadline"
	"github.com/Pivot-Studio/grpcdemo/interceptor/trace"

	"github.com/Pivot-Studio/grpcdemo/interceptor/ratelimit"

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
