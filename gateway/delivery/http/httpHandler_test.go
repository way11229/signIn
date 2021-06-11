package httpHandler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	httpHandler "signIn/gateway/delivery/http"
	"signIn/gateway/domain"
	"signIn/gateway/domain/mocks"
)

func TestGateway(t *testing.T) {
	accessData := domain.AccessData{
		Token: "7G7ovtjlalaCDzWtUVO2",
		Extra: "",
	}

	jsonExtra, _ := json.Marshal(domain.LineResponse{})
	expectResponseData := domain.SignInData{
		ID:    "123456789",
		Name:  "Way",
		Email: "test@test.com",
		Phone: "",
		Extra: string(jsonExtra),
	}

	rtnJson, _ := json.Marshal(gin.H{
		"message": "",
		"result":  expectResponseData,
	})

	mockSignInService := new(mocks.SignInService)
	mockSignInService.On(
		"SignInWithLine",
		mock.Anything,
		accessData,
	).Return(expectResponseData, nil).Once()

	r := gin.Default()

	handler := httpHandler.HttpHandler{
		SignInWithLineService: mockSignInService,
	}

	r.POST("/", handler.Gateway)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", strings.NewReader("method=line&verifyCode="+accessData.Token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(rtnJson), w.Body.String())
}