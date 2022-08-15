package interceptor

import "google.golang.org/grpc"

func GetUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		deadline,
	}
}
