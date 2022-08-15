package grpchelper

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// Get IP from GRPC context
func GetIP(ctx context.Context) string {
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		xForwardFor := headers.Get("x-forwarded-for")
		if len(xForwardFor) > 0 && xForwardFor[0] != "" {
			ips := strings.Split(xForwardFor[0], ",")
			if len(ips) > 0 {
				clientIp := ips[0]
				return clientIp
			}
		}
	}
	return ""
}
