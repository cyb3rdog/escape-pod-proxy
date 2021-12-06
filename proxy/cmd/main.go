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
	version  = "1.0.0"
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

func getEnvVars() {

	println("- Enviroment Variables:")
	println("CYB3RVECTOR_PROXY_CLIENT_PORT=" + os.Getenv("CYB3RVECTOR_PROXY_CLIENT_PORT"))
	println("CYB3RVECTOR_PROXY_CLIENT_CLIENT_AUTHENTICATION=" + os.Getenv("CYB3RVECTOR_PROXY_CLIENT_CLIENT_AUTHENTICATION"))
	println("CYB3RVECTOR_PROXY_CLIENT_INSECURE=" + os.Getenv("CYB3RVECTOR_PROXY_CLIENT_INSECURE"))
	println()
	println("CYB3RVECTOR_PROXY_SERVER_PORT=" + os.Getenv("CYB3RVECTOR_PROXY_SERVER_PORT"))
	println("CYB3RVECTOR_PROXY_SERVER_CLIENT_AUTHENTICATION=" + os.Getenv("CYB3RVECTOR_PROXY_SERVER_CLIENT_AUTHENTICATION"))
	println("CYB3RVECTOR_PROXY_SERVER_INSECURE=" + os.Getenv("CYB3RVECTOR_PROXY_SERVER_INSECURE"))
	println()
	println("CYB3RVECTOR_MONGO_DB_NAME=" + os.Getenv("CYB3RVECTOR_MONGO_DB_NAME"))
	println("CYB3RVECTOR_MONGO_DB_HOST=" + os.Getenv("CYB3RVECTOR_MONGO_DB_HOST"))
	println("CYB3RVECTOR_MONGO_DB_PORT=" + os.Getenv("CYB3RVECTOR_MONGO_DB_PORT"))
	println("CYB3RVECTOR_MONGO_DB_USER=" + os.Getenv("CYB3RVECTOR_MONGO_DB_USER"))
	println("CYB3RVECTOR_MONGO_DB_PASS=" + os.Getenv("CYB3RVECTOR_MONGO_DB_PASS"))
	println()
}

func printWin() {
	println()
	println(" --- ---------------------------------------------------------------------- ---")
	println(" --- In order for this proxy to be able to connect to EscapePod's mongo DB, ---")
	println(" --- MongoDB binding address needs to be changed in /etc/mongod.conf config ---")
	println(" --- file from localhost '127.0.0.1' to public '0.0.0.0'                    ---")
	println(" --- ---------------------------------------------------------------------- ---")
	println(" --- In order for the EscapePod itself to be able to connect to this proxy, ---")
	println(" --- and push its events here, it needs to know where this proxy is hosted. ---")
	println(" --- Set following variables in /etc/escape-pod.conf file corresondingly:   ---")
	println(" ---                                                                        ---")
	println(" --- ENABLE_EXTENSIONS=true                                                 ---")
	println(" --- ESCAPEPOD_EXTENDER_TARGET=XX.XX.XX.XX:8089                             ---")
	println(" --- ESCAPEPOD_EXTENDER_DISABLE_TLS=true                                    ---")
	println(" --- ---------------------------------------------------------------------- ---")
	println()
}

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")

	println("Cyb3rVector EscapePod Extension proxy v." + version)

	//setEnvVars()  // uncomment to hardcode the env vars
	getEnvVars()
	//printWin()

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
