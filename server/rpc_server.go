package server

import (
	"flag"
	"fmt"
	er "github.com/NineSwards/go-job/server/executor"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 3344, "The server port")
)

func StartRpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	er.RegisterJobExecutorRegisterServer(s, &er.RpcServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
