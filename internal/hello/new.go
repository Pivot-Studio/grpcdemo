package hello

import (
	"github.com/Pivot-Studio/grpcdemo/gen/hello_servicepb"
	"github.com/Pivot-Studio/grpcdemo/pkg/boot"

	"google.golang.org/grpc"
)

func RegisterHelloWordService(s *grpc.Server) {
	hello_servicepb.RegisterHelloServiceServer(s, &server{})
	boot.RegisterGatewayFunc(hello_servicepb.RegisterHelloServiceHandlerFromEndpoint)
}
