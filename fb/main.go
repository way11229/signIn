package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	grpcHandler "signIn/fb/delivery/grpc"
	"signIn/fb/domain"
	getAccessTokenRepo "signIn/fb/repository/getAccessToken"
	getUserProfileRepo "signIn/fb/repository/getUserProfile"
	VerifyTokenRepo "signIn/fb/repository/verifyToken"
	signInService "signIn/fb/service/signIn"
)

const GRPC_LISTEN_PORT = ":80"

var fbConfig domain.FbConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var hasRedirectUrl bool
	var hasClientId bool
	var hasClientSecret bool

	fbConfig.RedirectUrl, hasRedirectUrl = os.LookupEnv("REDIRECT_URL")
	if !hasRedirectUrl {
		panic("fb redirect url is empty")
	}

	fbConfig.ClientId, hasClientId = os.LookupEnv("CLIENT_ID")
	if !hasClientId {
		panic("fb client id is empty")
	}

	fbConfig.ClientSecret, hasClientSecret = os.LookupEnv("CLIENT_SECRET")
	if !hasClientSecret {
		panic("fb client secret is empty")
	}

	fbConfig.TokenApi = "https://graph.facebook.com/v11.0/oauth/access_token/"
	fbConfig.VerifyApi = "https://graph.facebook.com/debug_token/"
	fbConfig.ProfileApi = "https://graph.facebook.com/"
}

func main() {
	lis, err := net.Listen("tcp", GRPC_LISTEN_PORT)
	if err != nil {
		panic("net listen error")
	}

	signInServiceRepositoryList := domain.SignInServiceRepositoryList{
		GetAccessTokenRepository: getAccessTokenRepo.New(fbConfig),
		VerifyTokenRepository:    VerifyTokenRepo.New(fbConfig),
		GetUserProfileRepository: getUserProfileRepo.New(fbConfig),
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
