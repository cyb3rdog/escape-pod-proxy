
protoc -I./defs/podextension --csharp_out ./lang/csharp/podextension extension.proto
protoc -I./defs/podextension --grpc_out ./lang/csharp/podextension extension.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe
protoc -I./defs/podextension --go_out ./lang/go/podextension --go_opt=paths=source_relative extension.proto
protoc -I./defs/podextension --go-grpc_out ./lang/go/podextension --go-grpc_opt=paths=source_relative extension.proto


protoc -I./defs/cybervector --csharp_out ./lang/csharp/cybervector cybervector_proxy.proto
protoc -I./defs/cybervector --grpc_out ./lang/csharp/cybervector cybervector_proxy.proto --plugin=protoc-gen-grpc=grpc_csharp_plugin.exe
protoc -I./defs/cybervector --go_out ./lang/go/cybervector --go_opt=paths=source_relative cybervector_proxy.proto
protoc -I./defs/cybervector --go-grpc_out ./lang/go/cybervector --go-grpc_opt=paths=source_relative cybervector_proxy.proto
