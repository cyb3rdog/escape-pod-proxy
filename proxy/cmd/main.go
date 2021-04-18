// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// main - the proxy service entrypoint, starting all listenes and
// proxying them toghether.
//
package main

import (
	"os"

	"github.com/cyb3rdog/escape-pod-proxy/proxy/pkg/proxyclient"
	"github.com/cyb3rdog/escape-pod-proxy/proxy/pkg/proxyserver"
	"github.com/digital-dream-labs/hugh/grpc/server"
	"github.com/digital-dream-labs/hugh/log"
)

const (
	certFile = "../../../certgen/Cyb3rVector-Extension.crt"
	keyFile  = "../../../certgen/Cyb3rVector-Extension.pem"
)

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")

	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_PORT", "8089")
	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_CLIENT_AUTHENTICATION", "NoClientCert")
	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_INSECURE", "true")

	os.Setenv("CYB3RVECTOR_PROXY_SERVER_PORT", "8090")
	os.Setenv("CYB3RVECTOR_PROXY_SERVER_CLIENT_AUTHENTICATION", "NoClientCert")
	os.Setenv("CYB3RVECTOR_PROXY_SERVER_INSECURE", "true")

	proxyServer, err := proxyserver.New(
		proxyserver.WithViper("env-prefix", "CYB3RVECTOR_PROXY_SERVER"),
	)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer, err := proxyServer.GetGrpcServer()
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Start()

	proxyClient, err := proxyclient.New(
		proxyclient.WithViper("env-prefix", "CYB3RVECTOR_PROXY_CLIENT"),
		proxyclient.WithProxyTo(proxyServer),
	)
	if err != nil {
		log.Fatal(err)
	}
	grpcClient, err := proxyClient.GetGrpcServer()
	if err != nil {
		log.Fatal(err)
	}
	grpcClient.Start()

	<-grpcServer.Notify(server.Stopped)
	<-grpcClient.Notify(server.Stopped)
}
