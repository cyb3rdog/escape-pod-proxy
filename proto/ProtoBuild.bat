protoc -I./defs/podextension --go_out ./lang/go/podextension --go_opt=paths=source_relative extension.proto
protoc -I./defs/podextension --go-grpc_out ./lang/go/podextension --go-grpc_opt=paths=source_relative extension.proto
protoc -I./defs/cybervector --go_out ./lang/go/cybervector --go_opt=paths=source_relative cybervector_proxy.proto
protoc -I./defs/cybervector --go-grpc_out ./lang/go/cybervector --go-grpc_opt=paths=source_relative cybervector_proxy.proto

protoc -I./defs/podextension --csharp_out ./lang/csharp/podextension extension.proto
protoc -I./defs/podextension --grpc_out ./lang/csharp/podextension extension.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe
protoc -I./defs/cybervector --csharp_out ./lang/csharp/cybervector cybervector_proxy.proto
protoc -I./defs/cybervector --grpc_out ./lang/csharp/cybervector cybervector_proxy.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe

protoc -I./defs/podextension --python_out ./lang/python/podextension extension.proto
protoc -I./defs/podextension --grpclib_python_out ./lang/python/podextension extension.proto --plugin=grpc_python_plugin
protoc -I./defs/cybervector --python_out ./lang/python/cybervector cybervector_proxy.proto
protoc -I./defs/cybervector --grpclib_python_out ./lang/python/cybervector cybervector_proxy.proto --plugin=grpc_python_plugin

protoc -I./defs/podextension --js_out ./lang/javascript/podextension extension.proto
protoc -I./defs/cybervector --js_out ./lang/javascript/cybervector cybervector_proxy.proto
