package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"proto-test/gen/hello_servicepb"
	"proto-test/interceptor"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	hello_servicepb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *hello_servicepb.HelloRequest) (*hello_servicepb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello_servicepb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer(getServerOpts()...)
		hello_servicepb.RegisterHelloServiceServer(s, &server{})
		s.Serve(lis)
	}()
	opts := getDialOpts()
	gwmux := runtime.NewServeMux()
	hello_servicepb.RegisterHelloServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:9000", opts)

	fs := http.FileServer(http.Dir("./swaggerui"))
	h := http.StripPrefix("/swaggerui/", fs)
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/swaggerui/", h)
	http.ListenAndServe(":8080", mux)

}

func getServerOpts() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.GetUnaryInterceptors()...),
		grpc.ChainStreamInterceptor(),
	}
}

func getDialOpts() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}
