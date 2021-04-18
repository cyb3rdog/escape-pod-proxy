
protoc -I./podextension --csharp_out ./podextension extension.proto
protoc -I./podextension --grpc_out ./podextension extension.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe

protoc -I./cybervector --csharp_out ./cybervector cybervector_proxy.proto
protoc -I./cybervector --grpc_out ./cybervector cybervector_proxy.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe
