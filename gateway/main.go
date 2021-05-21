package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// httpHandler "signIn/gateway/delivery/http"
)

func main() {
	log.Info("Sign in gateway server start")

	r := gin.Default()
	// httpHandler.NewHttpHandler(r)
	r.LoadHTMLGlob("view/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	log.Fatal(r.Run(":80"))
}
