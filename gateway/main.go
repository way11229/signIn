package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"signIn/gateway/domain"
	cors "signIn/gateway/middleware/cors"

	httpHandler "signIn/gateway/delivery/http"
	lineRepository "signIn/gateway/repository/line"
	signInService "signIn/gateway/service/signIn"
)

const (
	GRPC_LINE_CONNECT = "line:80"
)

func main() {
	log.Info("Sign in gateway server start")

	r := gin.Default()
	r.Use(cors.CORSMiddleWare())

	lineGRPCConn := mGetLineGRPCConn()
	lr := lineRepository.New(lineGRPCConn)
	serviceList := domain.ServiceList{
		SignInService: signInService.New(lr),
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
