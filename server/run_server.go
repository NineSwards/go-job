package server

import (
	"flag"
	"fmt"
	reg "github.com/NineSwards/go-job/core/server/registry"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 3344, "The server port")
)

func RunServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reg.RegisterRegisterServer(s, &reg.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
