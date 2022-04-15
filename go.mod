module emqx.io/grpc/emqx_schema_registry

go 1.11

replace emqx.io/grpc/emqx_schema_registry => ./

require (
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0 // indirect
	google.golang.org/protobuf v1.28.0
)
