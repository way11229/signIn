package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"signIn/gateway/domain"
	cors "signIn/gateway/middleware/cors"

	httpHandler "signIn/gateway/delivery/http"

	signInService "signIn/gateway/service/signIn"

	fbRepository "signIn/gateway/repository/fb"
	googleRepository "signIn/gateway/repository/google"
	lineRepository "signIn/gateway/repository/line"
)

const (
	GRPC_LINE_CONNECT   = "signIn_line:80"
	GRPC_FB_CONNECT     = "signIn_fb:80"
	GRPC_GOOGLE_CONNECT = "signIn_google:80"
)

func main() {
	log.Info("Sign in gateway server start")

	r := gin.Default()
	r.Use(cors.CORSMiddleWare())

	lineGRPCConn := mGetLineGRPCConn()
	fbGRPCConn := mGetFbGRPCConn()
	googleGRPCConn := mGetGoogleGRPCConn()
	lr := lineRepository.New(lineGRPCConn)
	fr := fbRepository.New(fbGRPCConn)
	gr := googleRepository.New(googleGRPCConn)

	repositoryList := domain.RepositoryList{
		LineRepository:   lr,
		FbRepository:     fr,
		GoogleRepository: gr,
	}

	serviceList := domain.ServiceList{
		SignInService: signInService.New(repositoryList),
	}

	httpHandler.NewHttpHandler(r, serviceList)

	log.Fatal(r.Run(":80"))
}

func mGetLineGRPCConn() *grpc.ClientConn {
	grpcLineConnect, err := grpc.Dial(GRPC_LINE_CONNECT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("GRPC line connect error: " + err.Error())
		panic(err)
	}

	defer grpcLineConnect.Close()

	return grpcLineConnect
}

func mGetFbGRPCConn() *grpc.ClientConn {
	grpcFbConnect, err := grpc.Dial(GRPC_FB_CONNECT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("GRPC fb connect error: " + err.Error())
		panic(err)
	}

	defer grpcFbConnect.Close()

	return grpcFbConnect
}

func mGetGoogleGRPCConn() *grpc.ClientConn {
	grpcGoogleConnect, err := grpc.Dial(GRPC_GOOGLE_CONNECT, grpc.WithInsecure())
	if err != nil {
		log.Fatal("GRPC google connect error: " + err.Error())
		panic(err)
	}

	defer grpcGoogleConnect.Close()

	return grpcGoogleConnect
}
