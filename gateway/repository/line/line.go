package lineRepository

import (
	"context"
	"errors"
	"signIn/gateway/domain"
	ps "signIn/gateway/gen/line"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type signInWithLineRepository struct {
	grpcLineConnect *grpc.ClientConn
}

func New(cc *grpc.ClientConn) domain.LineRepository {
	return &signInWithLineRepository{grpcLineConnect: cc}
}

func (slr *signInWithLineRepository) SendSignInRequest(cxt context.Context, accessData domain.AccessData) (domain.LineResponse, error) {
	rtn := domain.LineResponse{}
	requestData := &ps.SignInData{
		VerifyCode: accessData.Token,
	}

	if slr.grpcLineConnect.GetState().String() == connectivity.Shutdown.String() {
		grpcLineConnect, gprcErr := grpc.Dial(slr.grpcLineConnect.Target(), grpc.WithInsecure())
		if gprcErr != nil {
			panic(gprcErr)
		}

		defer grpcLineConnect.Close()
		slr.grpcLineConnect = grpcLineConnect
	}

	lineSIgnInClient := ps.NewLineClient(slr.grpcLineConnect)
	grpcResponse, err := lineSIgnInClient.SignIn(cxt, requestData)
	if err != nil {
		return rtn, err
	}

	if grpcResponse.GetError() != "" {
		return rtn, errors.New(grpcResponse.GetError())
	}

	rtn = domain.LineResponse{
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
