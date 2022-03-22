package server

import (
	"golang.org/x/net/context"
	"log"
)

type RpcServer struct {
	UnimplementedJobExecutorRegisterServer
}

func (s *RpcServer) ExecutorRegistry(_ context.Context, in *RegistryRequest) (*RegistryReply, error) {

	log.Printf("Received: %v", in.GetServerName())
	serverList := in.GetServerList()
	for idx, value := range serverList {
		log.Printf("Received: index %d server %s", idx, value)
	}
	log.Printf("Received serverList: %v", serverList)
	return &RegistryReply{
		Message: "Hello " + in.GetServerName(),
	}, nil
}

func (s *RpcServer) ExecutorUnRegistry(_ context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetServerName())
	return &RegistryReply{
		Message: "Bye " + in.GetServerName(),
	}, nil
}
