package getAccessTokenRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"signIn/line/domain"
)

type getAccessTokenRepository struct {
	LineConfig domain.LineConfig
}

func New(lc domain.LineConfig) domain.GetAccessTokenRepository {
	return &getAccessTokenRepository{LineConfig: lc}
}

func (g *getAccessTokenRepository) GetAccessToken(verifyCode string) (domain.AccessTokenResponse, error) {
	rtn := domain.AccessTokenResponse{}

	requestBody := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {verifyCode},
		"redirect_uri":  {g.LineConfig.RedirectUrl},
		"client_id":     {g.LineConfig.ClientId},
		"client_secret": {g.LineConfig.ClientSecret},
	}

	response, err := http.PostForm(g.LineConfig.TokenApi, requestBody)
	if err != nil {
		return rtn, err
	}

	defer response.Body.Close()

	var responseDecode domain.AccessTokenResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseDecode)
	if decodeErr != nil {
		return rtn, decodeErr
	}

	if responseDecode.Error != "" {
		return rtn, errors.New(responseDecode.ErrorDescription)
	}

	rtn = responseDecode

	return rtn, nil
}
