package httpHandler

import (
	"net/http"
	"signIn/gateway/domain"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type HttpHandler struct {
	SignInService domain.SignInService
}

func NewHttpHandler(e *gin.Engine, serviceList domain.ServiceList) {
	handler := &HttpHandler{
		SignInService: serviceList.SignInService,
	}

	e.POST("/", handler.Gateway)
}

func (h *HttpHandler) Gateway(c *gin.Context) {
	var err error
	var signInData domain.SignInData

	method := c.DefaultPostForm("method", "")
	verifyCode := c.DefaultPostForm("verifyCode", "")
	if method == "" || verifyCode == "" {
		log.Warn("Sign In error: method:" + method + ",error: Missing paramters")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing paramters",
			"result":  signInData,
		})
		return
	}

	accessData := domain.AccessData{
		Token: verifyCode,
		Extra: c.DefaultPostForm("extra", ""),
	}

	switch method {
	case domain.LINE_METHOD:
		signInData, err = h.SignInService.SignInWithLine(c, accessData)
	case domain.FB_METHOD:
		signInData, err = h.SignInService.SignInWithFb(c, accessData)
	case domain.GOOGLE_METHOD:
		signInData, err = h.SignInService.SignInWithGoogle(c, accessData)
	}

	if err != nil {
		message := err.Error()
		log.Warn("Sign In error: method:" + method + ",error:" + message)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"result":  signInData,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"result":  signInData,
	})
}
