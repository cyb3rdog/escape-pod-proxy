// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// proxyclient - handles the incomming connection from the escapepod binary
//
package proxyclient

import (
	"fmt"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension"
	"github.com/digital-dream-labs/hugh/grpc/interceptors/log"
	"github.com/digital-dream-labs/hugh/grpc/server"

	logger "github.com/digital-dream-labs/hugh/log"
)

// ProxyClient is the configuration struct
type ProxyClient struct {
	podextension.UnimplementedPodExtensionServer
	server *server.Server
	proxy  ProxyTarget
}

// New returns a populated ProxyClient struct
func New(opts ...Option) (*ProxyClient, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.proxy == nil {
		return nil, fmt.Errorf("no target proxy server defined")
	}

	env_prefix := "PROXY_CLIENT"
	if cfg.prefix != "" {
		env_prefix = cfg.prefix
	}

	srv, err := server.New(
		server.WithViper("env-prefix", env_prefix),
		server.WithLogger(logger.Base()),
		server.WithReflectionService(),

		server.WithUnaryServerInterceptors(
			log.UnaryServerInterceptor(),
		),

		server.WithStreamServerInterceptors(
			log.StreamServerInterceptor(),
		),
	)
	if err != nil {
		logger.Fatal(err)
	}

	proxy := ProxyClient{
		server: srv,
		proxy:  cfg.proxy,
	}

	podextension.RegisterPodExtensionServer(srv.Transport(), &proxy)

	return &proxy, nil
}

// GetGrpcServer returns the internal grpc Server object
func (client *ProxyClient) GetGrpcServer() (*server.Server, error) {
	if client.server == nil {
		return nil, fmt.Errorf("no grpc server defined")
	}
	return client.server, nil
}
