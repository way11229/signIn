package httpHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	// SignInWithLineService domain.SignInWithLineService
}

func NewHttpHandler(e *gin.Engine) {
	handler := &HttpHandler{
		// SignInWithLineService: serviceList.SignInWithLineService,
	}

	e.POST("/line", handler.LineSignIn)
}

func (h *HttpHandler) LineSignIn(c *gin.Context) {
	verifyCode := c.PostForm("verifyCode")
	if verifyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "signin-frontend.selldarity.com")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

	c.JSON(http.StatusOK, gin.H{
		"code": verifyCode,
	})
}
