package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	grpcHandler "signIn/line/delivery/grpc"
	"signIn/line/domain"
	getAccessTokenRepo "signIn/line/repository/getAccessToken"
	getUserDataRepo "signIn/line/repository/getUserData"
	signInService "signIn/line/service/signIn"
)

const GRPC_LISTEN_PORT = ":80"

var lineConfig domain.LineConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var hasRedirectUrl bool
	var hasClientId bool
	var hasClientSecret bool

	lineConfig.RedirectUrl, hasRedirectUrl = os.LookupEnv("REDIRECT_URL")
	if !hasRedirectUrl {
		panic("line redirect url is empty")
	}

	lineConfig.ClientId, hasClientId = os.LookupEnv("CLIENT_ID")
	if !hasClientId {
		panic("line client id is empty")
	}

	lineConfig.ClientSecret, hasClientSecret = os.LookupEnv("CLIENT_SECRET")
	if !hasClientSecret {
		panic("line client secret is empty")
	}

	lineConfig.TokenApi = "https://api.line.me/oauth2/v2.1/token"
	lineConfig.VerifyApi = "https://api.line.me/oauth2/v2.1/verify"
}

func main() {
	lis, err := net.Listen("tcp", GRPC_LISTEN_PORT)
	if err != nil {
		panic("net listen error")
	}

	signInServiceRepositoryList := domain.SignInServiceRepositoryList{
		GetAccessTokenRepository: getAccessTokenRepo.New(lineConfig),
		GetUserDataRepository:    getUserDataRepo.New(lineConfig),
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
