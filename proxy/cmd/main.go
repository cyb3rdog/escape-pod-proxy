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

	"github.com/cyb3rdog/escape-pod-proxy/proxy/pkg/mongoclient"
	"github.com/cyb3rdog/escape-pod-proxy/proxy/pkg/proxyclient"
	"github.com/cyb3rdog/escape-pod-proxy/proxy/pkg/proxyserver"
	"github.com/digital-dream-labs/hugh/grpc/server"
	"github.com/digital-dream-labs/hugh/log"
)

const (
	certFile = "../../../certgen/Cyb3rVector-Extension.crt"
	keyFile  = "../../../certgen/Cyb3rVector-Extension.pem"
)

func setEnvVars() {

	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_PORT", "8089")
	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_CLIENT_AUTHENTICATION", "NoClientCert")
	os.Setenv("CYB3RVECTOR_PROXY_CLIENT_INSECURE", "true")

	os.Setenv("CYB3RVECTOR_PROXY_SERVER_PORT", "8090")
	os.Setenv("CYB3RVECTOR_PROXY_SERVER_CLIENT_AUTHENTICATION", "NoClientCert")
	os.Setenv("CYB3RVECTOR_PROXY_SERVER_INSECURE", "true")

	os.Setenv("CYB3RVECTOR_MONGO_DB_NAME", "database")
	os.Setenv("CYB3RVECTOR_MONGO_DB_HOST", "escapepod.local")
	os.Setenv("CYB3RVECTOR_MONGO_DB_PORT", "27017")
	os.Setenv("CYB3RVECTOR_MONGO_DB_USER", "myUserAdmin")
	os.Setenv("CYB3RVECTOR_MONGO_DB_PASS", "MzBmMWFmY2NhYzE0")

}

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")

	// uncomment to hardcode the env vars
	setEnvVars()

	mongoClient, err := mongoclient.New(
		mongoclient.WithViper("env-prefix", "CYB3RVECTOR_MONGO_DB"),
	)
	if err != nil {
		log.Fatal(err)
	}

	proxyServer, err := proxyserver.New(
		proxyserver.WithViper("env-prefix", "CYB3RVECTOR_PROXY_SERVER"),
		proxyserver.WithIntentFactory(mongoClient),
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
