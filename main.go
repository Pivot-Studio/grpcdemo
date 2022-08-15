package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/Pivot-Studio/grpcdemo/interceptor"
	"github.com/Pivot-Studio/grpcdemo/internal/hello"
	"github.com/Pivot-Studio/grpcdemo/pkg/boot"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("panic: %v", err)
		}
	}()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(getServerOpts()...)
	registerGRPC(s)
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	opts := getDialOpts()
	gwmux := runtime.NewServeMux()
	err = boot.RegisterGWs(context.Background(), gwmux, fmt.Sprintf("localhost:%d", 9000), opts)
	if err != nil {
		log.Fatalf("failed to register ep: %v", err)
	}
	fs := http.FileServer(http.Dir("./swaggerui"))
	h := http.StripPrefix("/swaggerui/", fs)
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/swaggerui/", h)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerGRPC(s *grpc.Server) {
	hello.RegisterHelloWordService(s)
}

func getServerOpts() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.GetUnaryInterceptors()...),
		grpc.ChainStreamInterceptor(),
		grpc.InTapHandle(interceptor.GetTapHandler()),
	}
}

func getDialOpts() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}
