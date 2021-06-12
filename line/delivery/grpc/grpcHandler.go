package grpcHandler

import (
	"context"
	pb "signIn/line/gen/line"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedLineServer
}

func New(s *grpcLib.Server) {
	handler := &GrpcHandler{}

	pb.RegisterLineServer(s, handler)
}

func (g *GrpcHandler) SignIn(cxt context.Context, accessData *pb.AccessData) (*pb.LineResponse, error) {
	rtn := pb.LineResponse{}
	if accessData.Token == "" {
		rtn.Error = "Token is missing"
		return &rtn, nil
	}

	return &rtn, nil
}
