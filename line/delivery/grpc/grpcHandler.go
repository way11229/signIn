package grpcHandler

import (
	"context"
	"signIn/line/domain"
	pb "signIn/line/gen/line"

	grpcLib "google.golang.org/grpc"
)

type GrpcHandler struct {
	pb.UnimplementedLineServer

	SignInService domain.SignInService
}

func New(s *grpcLib.Server, serviceList domain.ServiceList) {
	handler := &GrpcHandler{
		SignInService: serviceList.SignInService,
	}

	pb.RegisterLineServer(s, handler)
}

func (g *GrpcHandler) SignIn(cxt context.Context, signInData *pb.SignInData) (*pb.LineResponse, error) {
	rtn := pb.LineResponse{}
	if signInData.VerifyCode == "" {
		rtn.Error = "Verify code is empty"
		return &rtn, nil
	}

	signInResponse, err := g.SignInService.SignIn(signInData.VerifyCode)
	if err != nil {
		rtn.Error = err.Error()
		return &rtn, nil
	}

	rtn = pb.LineResponse{
		AccessToken:         signInResponse.AccessToken,
		AccessTokenExpireIn: signInResponse.AccessTokenExpireIn,
		RefreshToken:        signInResponse.RefreshToken,
		UserId:              signInResponse.UserId,
		Name:                signInResponse.Name,
		Picture:             signInResponse.Picture,
		Email:               signInResponse.Email,
		StatusMessage:       signInResponse.StatusMessage,
		Error:               "",
	}

	return &rtn, nil
}
