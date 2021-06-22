package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	grpcHandler "signIn/google/delivery/grpc"
	"signIn/google/domain"
	getAccessTokenRepo "signIn/google/repository/getAccessToken"
	getUserInfoRepo "signIn/google/repository/getUserInfo"
	signInService "signIn/google/service/signIn"
)

const GRPC_LISTEN_PORT = ":80"

var googleConfig domain.GoogleConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var hasRedirectUrl bool
	var hasClientId bool
	var hasClientSecret bool

	googleConfig.RedirectUrl, hasRedirectUrl = os.LookupEnv("REDIRECT_URL")
	if !hasRedirectUrl {
		panic("google redirect url is empty")
	}

	googleConfig.ClientId, hasClientId = os.LookupEnv("CLIENT_ID")
	if !hasClientId {
		panic("google client id is empty")
	}

	googleConfig.ClientSecret, hasClientSecret = os.LookupEnv("CLIENT_SECRET")
	if !hasClientSecret {
		panic("google client secret is empty")
	}

	googleConfig.TokenApi = "https://oauth2.googleapis.com/token"
	googleConfig.UserInfoApi = "https://www.googleapis.com/oauth2/v3/userinfo"
}

func main() {
	lis, err := net.Listen("tcp", GRPC_LISTEN_PORT)
	if err != nil {
		panic("net listen error")
	}

	signInServiceRepositoryList := domain.SignInServiceRepositoryList{
		GetAccessTokenRepository: getAccessTokenRepo.New(googleConfig),
		GetUserInfoRepository:    getUserInfoRepo.New(googleConfig),
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
