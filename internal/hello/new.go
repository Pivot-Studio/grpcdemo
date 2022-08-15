package hello

import (
	"proto-test/gen/hello_servicepb"
	"proto-test/pkg/boot"

	"google.golang.org/grpc"
)

func RegisterHelloWordService(s *grpc.Server) {
	hello_servicepb.RegisterHelloServiceServer(s, &server{})
	boot.RegisterGatewayFunc(hello_servicepb.RegisterHelloServiceHandlerFromEndpoint)
}
