package httpHandler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type HttpHandler struct {
	// SignInWithLineService domain.SignInWithLineService
}

func NewHttpHandler(e *gin.Engine) {
	handler := &HttpHandler{
		// SignInWithLineService: serviceList.SignInWithLineService,
	}

	e.LoadHTMLGlob("view/*")
	e.GET("/", handler.ShowFrontEnd)
}

func (h *HttpHandler) ShowFrontEnd(c *gin.Context) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	c.JSON(200, gin.H{
		"test": os.Getenv("SIGN_IN_GATEWAY_HOST"),
	})
}
