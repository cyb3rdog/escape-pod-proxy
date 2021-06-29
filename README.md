
# Cyb3rVector Escape-Pod Proxy

This Proxy has beed developed as a result of non-suitable architecture design of the escape-pod and its extensions ([this issue](https://github.com/digital-dream-labs/escape-pod-extension/issues/4)). It resolves the problem of need for having multiple clients/extensions conneted to the same escape pod, by creating an additional communication layer, implemented as server-side event streaming service which fires the intent events generated by the escape-pod to all of its subscribed client.
Additionally, this proxy also contains an gRPC interface for realtime manipulation of the Escapepod MongoDB intent database, to satisfy the need for dynamic creattion of the intent definitions.

This proxy has been implemented in the [Cyb3rVector CodeLab](https://cyb3rdog.github.io/Cyb3rVector) application, where the blockly code programs capabilities has been extended with ability to dynamically create an intents with given keywords and phrases and to react to them in the blockly code, by handing the events received from this proxy.

It's use however is not limited to this application and can be used anywhere, wherever there is a need for reacting or altering the `escapepod extend intents` client-side.


## General information 

This repository contains two main branches:
1) [MASTER](https://github.com/cyb3rdog/escape-pod-proxy/tree/master) - This Extension Proxy, written in go-lang
2) [Cyb3rPod](https://github.com/cyb3rdog/escape-pod-proxy/tree/Cyb3rPod) - Default simple client implementation in c#


### Original architecture

This diagram describes the architecture of the orginal escapepod extensibility implementation and its flaws:

![image](https://user-images.githubusercontent.com/12493945/123772149-e9f46200-d8cb-11eb-8cc8-c1388800878e.png)

### Cyb3rVector Extension Proxy architecture

Following diagram describes the difference of how the escapepod extensibility has been implemented for the needs of the Cyb3rVector, where the unary grpc server calls are transposed into the gRPC event stream:

![image](https://user-images.githubusercontent.com/12493945/123773145-bbc35200-d8cc-11eb-9608-cfd045f99bb1.png)


## Content

This proxy consists of three main submodules:

1) [ProxyClient](https://github.com/cyb3rdog/escape-pod-proxy/blob/master/proxy/pkg/proxyclient) - gRPC server handling incomming connection from the Escapepod
2) [ProxyServer](https://github.com/cyb3rdog/escape-pod-proxy/blob/master/proxy/pkg/proxyserver) - gRPC server handling connection from clients, notifying them about intent events.
3) [MongoClient](https://github.com/cyb3rdog/escape-pod-proxy/tree/master/proxy/pkg/mongoclient) - MongoDB client allowing realtime intent database manipulation

***Note:***
*Escapepod binary itself surelly does already contain the interface for the mongoDB intent database manipulation (creating/editing/deleting intents), as this functionality is available in the Escapepod WebUI, but because of the lack of the techincal documentation, it was much easier and quicker to just implement the own custom MongoClient.*


## Build

This section describes how to build this extension proxy

1. Export the related environment variables

|Variable| Value | Description |
|--|--|--|
| CYB3RVECTOR_PROXY_CLIENT_PORT | 8089 | the TCP port the cybervector-proxy connects to escape-pod |
| CYB3RVECTOR_PROXY_CLIENT_INSECURE | true | Insecure TLS-free GRPC communication |
| CYB3RVECTOR_PROXY_CLIENT_CLIENT_AUTHENTICATION | NoClientCert | Insecure TLS-free GRPC communication |
| CYB3RVECTOR_PROXY_SERVER_PORT | 8090 | the TCP port the cybervector-proxy publishes events to clients |
| CYB3RVECTOR_PROXY_SERVER_INSECURE | true | Insecure TLS-free GRPC communication |
| CYB3RVECTOR_PROXY_SERVER_CLIENT_AUTHENTICATION | NoClientCert | Insecure TLS-free GRPC communication |

2. Build the ```cybervector-proxy``` service binary file

- Linux: ```$ make build```
- Windows: ```build.cmd```


## Deployment

#### A) Automatic, from the Cyb3rVector application:

The [Cyb3rVector](https://cyb3rdog.github.io/Cyb3rVector) application contains an automatic mechanism for deployment of this extension proxy to the EscapePod server.
For the insight how this deployment works, check the `sh` script in the [Cyb3rPod branch](https://github.com/cyb3rdog/escape-pod-proxy/tree/Cyb3rPod/Resources).

#### B) Manual:

This section describes how to deploy the Cyb3rVector Escape-Pod Proxy service to your Escape-Pod.

1. Make sure you've added the following lines to your escape pods config (in /etc/escape-pod.conf)

```sh
ENABLE_EXTENSIONS=true
ESCAPEPOD_EXTENDER_TARGET=127.0.0.1:8089
ESCAPEPOD_EXTENDER_DISABLE_TLS=true
```

2. Deploy the ```cybervector-proxy``` binary to your escape pod (i.e */usr/local/escapepod/bin/*)
4. Create service, to run the service during EscapePod boot with all the enviroment varibles initialized ([example service here](https://github.com/cyb3rdog/escape-pod-proxy/blob/Cyb3rPod/Resources/cybervector-proxy.service))
5. Restart the services / or your Escape-Pod
6. Try the [Cyb3rPod client](https://github.com/cyb3rdog/escape-pod-proxy/tree/Cyb3rPod/Resources) or [Cyb3rVector](https://cyb3rdog.github.io/Cyb3rVector) application!


## Testing

You can test your build either locally, from the enviroment you develope in, or from the docker.

1. IP Address

For any case, you need to know the IP address of the machine the extension service will run on, and 
modify the ```ESCAPEPOD_EXTENDER_TARGET``` variable in your ```/etc/escape-pod.conf``` file correspondingly.

2. Start your application!


### Enjoy!