package main

import (
	"net"

	"google.golang.org/grpc"

	grpcHandler "signIn/fb/delivery/grpc"
	"signIn/fb/domain"
	getUserProfileRepo "signIn/fb/repository/getUserProfile"
	signInService "signIn/fb/service/signIn"
)

const GRPC_LISTEN_PORT = ":80"

var fbConfig domain.FbConfig

func init() {
	fbConfig.GraphApi = "https://graph.facebook.com/"
}

func main() {
	lis, err := net.Listen("tcp", GRPC_LISTEN_PORT)
	if err != nil {
		panic("net listen error")
	}

	signInServiceRepositoryList := domain.SignInServiceRepositoryList{
		GetUserProfileRepository: getUserProfileRepo.New(fbConfig),
	}

	ServiceList := domain.ServiceList{
		SignInService: signInService.New(signInServiceRepositoryList),
	}

	s := grpc.NewServer()
	grpcHandler.New(s, ServiceList)

	if err := s.Serve(lis); err != nil {
		panic("grpc server error")
	}
}
