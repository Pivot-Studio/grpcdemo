package hello

import (
	"context"
	"log"
	"proto-test/gen/hello_servicepb"
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
