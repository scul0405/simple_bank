package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader = "user-agent"
	xForwarderForHeader = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata{
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		log.Printf("metadata: %v", md)

		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}

		if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}

		if clientIP := md.Get(xForwarderForHeader); len(clientIP) > 0 {
			mtdt.ClientIP = clientIP[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}