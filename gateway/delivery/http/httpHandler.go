package httpHandler

import (
	"net/http"
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

	e.LoadHTMLGlob("../../view/*")
	e.GET("/", handler.ShowFrontEnd)
}

func (h *HttpHandler) ShowFrontEnd(c *gin.Context) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"template": os.Getenv("SIGN_IN_GATEWAY_HOST"),
	})
}
