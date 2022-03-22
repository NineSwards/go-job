# go-job

## RPC service

```shell
brew install protobuf
```
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 

 protoc --java_out=./server/executor --go-grpc_out=./server/executor server/proto/executor_registry.proto
```