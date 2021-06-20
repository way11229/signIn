package fbRepository

import (
	"context"
	"errors"
	"signIn/gateway/domain"
	ps "signIn/gateway/gen/fb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type signInWithFbRepository struct {
	grpcFbConnect *grpc.ClientConn
}

func New(cc *grpc.ClientConn) domain.FbRepository {
	return &signInWithFbRepository{grpcFbConnect: cc}
}

func (slr *signInWithFbRepository) SendSignInRequest(cxt context.Context, accessData domain.AccessData) (domain.FbResponse, error) {
	rtn := domain.FbResponse{}
	requestData := &ps.SignInData{
		VerifyCode: accessData.Token,
	}

	if slr.grpcFbConnect.GetState().String() == connectivity.Shutdown.String() {
		grpcLineConnect, gprcErr := grpc.Dial(slr.grpcFbConnect.Target(), grpc.WithInsecure())
		if gprcErr != nil {
			panic(gprcErr)
		}

		defer grpcLineConnect.Close()
		slr.grpcFbConnect = grpcLineConnect
	}

	fbSIgnInClient := ps.NewFbClient(slr.grpcFbConnect)
	grpcResponse, err := fbSIgnInClient.SignIn(cxt, requestData)
	if err != nil {
		return rtn, err
	}

	if grpcResponse.GetError() != "" {
		return rtn, errors.New(grpcResponse.GetError())
	}

	rtn = domain.FbResponse{
		AccessToken:         grpcResponse.GetAccessToken(),
		AccessTokenExpireIn: grpcResponse.GetAccessTokenExpireIn(),
		UserId:              grpcResponse.GetUserId(),
		Name:                grpcResponse.GetName(),
		Picture:             grpcResponse.GetPicture(),
		Email:               grpcResponse.GetEmail(),
		Birthday:            grpcResponse.GetBirthday(),
	}

	return rtn, nil
}
