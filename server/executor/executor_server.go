package server

import (
	"golang.org/x/net/context"
	"log"
)

type RpcServer struct {
	UnimplementedJobExecutorRegisterServer
}

func (s *RpcServer) ExecutorRegistry(ctx context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetServerName())
	return &RegistryReply{Message: "Hello " + in.GetServerName()}, nil
}

func (s *RpcServer) ExecutorUnRegistry(ctx context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetServerName())
	return &RegistryReply{Message: "Bye " + in.GetServerName()}, nil
}
