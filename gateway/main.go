package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	httpHandler "signIn/gateway/delivery/http"

	cors "signIn/gateway/middleware/cors"
)

func main() {
	log.Info("Sign in gateway server start")

	r := gin.Default()
	r.Use(cors.CORSMiddleWare())
	httpHandler.NewHttpHandler(r)

	log.Fatal(r.Run(":80"))
}
