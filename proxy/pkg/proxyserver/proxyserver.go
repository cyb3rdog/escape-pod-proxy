// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// proxyserver - handles the incomming connection from the clients and
// notifies them about the intent events raised from the escape-pod.
//
package proxyserver

import (
	"fmt"

	"github.com/digital-dream-labs/hugh/grpc/interceptors/log"
	"github.com/digital-dream-labs/hugh/grpc/server"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/cybervector"
	interfaces "github.com/cyb3rdog/escape-pod-proxy/proxy/pkg"

	logger "github.com/digital-dream-labs/hugh/log"
)

type Subscriber struct {
	done   chan bool
	stream cybervector.CyberVectorProxyService_SubscribeServer
}

// ProxyServer is the configuration struct
type ProxyServer struct {
	cybervector.UnimplementedCyberVectorProxyServiceServer
	factory     interfaces.IntentFactory
	server      *server.Server
	subscribers map[string]Subscriber
}

// New returns a populated ProxyServer struct
func New(opts ...Option) (*ProxyServer, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}

	env_prefix := "PROXY_SERVER"
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

	proxy := ProxyServer{
		server:      srv,
		factory:     cfg.factory,
		subscribers: make(map[string]Subscriber),
	}

	cybervector.RegisterCyberVectorProxyServiceServer(srv.Transport(), &proxy)

	return &proxy, nil
}

// GetGrpcServer returns the internal grpc Server object
func (client *ProxyServer) GetGrpcServer() (*server.Server, error) {
	if client.server == nil {
		return nil, fmt.Errorf("no grpc server defined")
	}
	return client.server, nil
}
