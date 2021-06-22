package googleRepository

import (
	"context"
	"errors"
	"signIn/gateway/domain"
	ps "signIn/gateway/gen/google"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type signInWithGoogleRepository struct {
	grpcGoogleConnect *grpc.ClientConn
}

func New(cc *grpc.ClientConn) domain.GoogleRepository {
	return &signInWithGoogleRepository{grpcGoogleConnect: cc}
}

func (slr *signInWithGoogleRepository) SendSignInRequest(cxt context.Context, accessData domain.AccessData) (domain.GoogleResponse, error) {
	rtn := domain.GoogleResponse{}
	requestData := &ps.SignInData{
		VerifyCode: accessData.Token,
	}

	if slr.grpcGoogleConnect.GetState().String() == connectivity.Shutdown.String() {
		grpcLineConnect, gprcErr := grpc.Dial(slr.grpcGoogleConnect.Target(), grpc.WithInsecure())
		if gprcErr != nil {
			panic(gprcErr)
		}

		defer grpcLineConnect.Close()
		slr.grpcGoogleConnect = grpcLineConnect
	}

	fbSIgnInClient := ps.NewGoogleClient(slr.grpcGoogleConnect)
	grpcResponse, err := fbSIgnInClient.SignIn(cxt, requestData)
	if err != nil {
		return rtn, err
	}

	if grpcResponse.GetError() != "" {
		return rtn, errors.New(grpcResponse.GetError())
	}

	rtn = domain.GoogleResponse{
		AccessToken:         grpcResponse.GetAccessToken(),
		AccessTokenExpireIn: grpcResponse.GetAccessTokenExpireIn(),
		RefreshToken:        grpcResponse.GetRefreshToken(),
		UserId:              grpcResponse.GetUserId(),
		Name:                grpcResponse.GetName(),
		Picture:             grpcResponse.GetPicture(),
		Email:               grpcResponse.GetEmail(),
	}

	return rtn, nil
}
