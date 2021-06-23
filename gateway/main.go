package main

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"signIn/gateway/domain"
	cors "signIn/gateway/middleware/cors"

	grpcHandler "signIn/gateway/delivery/grpc"
	httpHandler "signIn/gateway/delivery/http"

	signInService "signIn/gateway/service/signIn"

	fbRepository "signIn/gateway/repository/fb"
	googleRepository "signIn/gateway/repository/google"
	lineRepository "signIn/gateway/repository/line"
)

const (
	GATEWAY_PORT        = ":80"
	GRPC_LINE_CONNECT   = "signIn_line:80"
	GRPC_FB_CONNECT     = "signIn_fb:80"
	GRPC_GOOGLE_CONNECT = "signIn_google:80"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	log.Info("Sign in gateway server start")

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

	mode, hasMode := os.LookupEnv("CONNECT_MODE")
	if hasMode && (mode == "grpc") {
		lis, listenErr := net.Listen("tcp", GATEWAY_PORT)
		if listenErr != nil {
			panic("net listen error")
		}
		s := grpc.NewServer()
		grpcHandler.New(s, serviceList)

		log.Fatal(s.Serve(lis))
	} else {
		r := gin.Default()
		r.Use(cors.CORSMiddleWare())
		httpHandler.NewHttpHandler(r, serviceList)

		log.Fatal(r.Run(GATEWAY_PORT))
	}
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
