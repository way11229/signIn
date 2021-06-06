package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	grpcHandler "signIn/line/delivery/grpc"
)

const GRPC_LISTEN_PORT = ":80"

func main() {
	lis, err := net.Listen("tcp", GRPC_LISTEN_PORT)
	if err != nil {
		fmt.Println("net listen error")
	}

	s := grpc.NewServer()
	grpcHandler.New(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println("grpc server error")
	}
}
