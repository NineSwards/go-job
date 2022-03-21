package registry

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	UnimplementedRegisterServer
}

func (s *Server) Registry(ctx context.Context, in *RegistryRequest) (*RegistryReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &RegistryReply{Message: "Hello " + in.GetName()}, nil
}
