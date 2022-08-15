package boot

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GwFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

var gatewayFuncs = []GwFunc{}

func RegisterGatewayFunc(f GwFunc) {
	gatewayFuncs = append(gatewayFuncs, f)
}

func RegisterGWs(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	for _, f := range gatewayFuncs {
		err = f(ctx, mux, endpoint, opts)
		if err != nil {
			return
		}
	}
	return

}
