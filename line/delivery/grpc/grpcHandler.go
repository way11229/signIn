package grpcHandler

import (
	"context"
	pb "signIn/line/gen/lineSignIn"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedLineSignInServer
}

func New(s *grpcLib.Server) {
	handler := &GrpcHandler{}

	pb.RegisterLineSignInServer(s, handler)
}

func (g *GrpcHandler) Query(cxt context.Context, accessData *pb.AccessData) (*pb.LineResponse, error) {
	rtn := pb.LineResponse{
		AccessToken:         "123",
		AccessTokenExpireIn: 1123456789,
		RefreshToken:        "abcdefg",
		UserId:              "abcdefg",
		Name:                "test",
		Picture:             "",
		Email:               "Email",
		Error:               "",
	}
	return &rtn, nil
}
