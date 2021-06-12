package lineRepository

import (
	"context"
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
	requestData := &ps.AccessData{
		Token: accessData.Token,
		Extra: accessData.Extra,
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
