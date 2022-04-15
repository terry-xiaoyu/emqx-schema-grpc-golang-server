# emqx_schema_registry-svr-go

This is a demo server written in golang for emqx_schema_registry

## Prerequisites

- [Go](https://golang.org) (any one of the three latest major)
- [Protocol buffer](https://developers.google.com/protocol-buffers) **compiler**, `protoc`
    For installation instructions, see
    [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).
- **Go plugins** for the protocol compiler:
    - Install the protocol compiler plugins for Go using the following commands:
    ```
    export GO111MODULE=on  # Enable module mode
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
           google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

    - Update your PATH so that the protoc compiler can find the plugins:
    ```
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## Run

Try to compile the `*.proto` files:

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuf/emqx_schema_registry.proto
```

Run server
```
go run main.go
```
