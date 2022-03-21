package registry

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	UnimplementedJobServerRegisterServer
}

func (s *Server) JobRegistry(ctx context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetServerName())
	return &RegistryReply{Message: "Hello " + in.GetServerName()}, nil
}

func (s *Server) JobUnRegistry(ctx context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetServerName())
	return &RegistryReply{Message: "Bye " + in.GetServerName()}, nil
}
