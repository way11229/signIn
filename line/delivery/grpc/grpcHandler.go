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

func (g *GrpcHandler) SignIn(cxt context.Context, signInData *pb.SignInData) (*pb.LineResponse, error) {
	rtn := pb.LineResponse{}
	if signInData.VerifyCode == "" {
		rtn.Error = "Verify code is empty"
		return &rtn, nil
	}

	return &rtn, nil
}
