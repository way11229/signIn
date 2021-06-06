package httpHandler

import (
	"net/http"
	"signIn/gateway/domain"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

const LINE_METHOD = "line"

type HttpHandler struct {
	SignInWithLineService domain.SignInService
}

func NewHttpHandler(e *gin.Engine, serviceList domain.ServiceList) {
	handler := &HttpHandler{
		SignInWithLineService: serviceList.SignInService,
	}

	e.POST("/", handler.Gateway)
}

func (h *HttpHandler) Gateway(c *gin.Context) {
	method := c.DefaultPostForm("method", "")
	verifyCode := c.DefaultPostForm("verifyCode", "")
	if method == "" || verifyCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	accessData := domain.AccessData{
		Token: verifyCode,
		Extra: "",
	}

	var err error
	var signInData domain.SignInData
	message := ""

	switch method {
	case LINE_METHOD:
		signInData, err = h.SignInWithLineService.SignInWithLine(c, accessData)
		break
	}

	if err != nil {
		message = err.Error()
		log.Fatal("Sign In error: method:" + method + ",error:" + message)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
		"result":  signInData,
	})
}
