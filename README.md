# go-job

## RPC service

```shell
brew install protobuf
```
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 

protoc --go_out=./server/registry --go-grpc_out=./server/registry server/registry/job_registry.proto
```