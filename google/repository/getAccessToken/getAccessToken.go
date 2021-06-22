package getAccessTokenRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"signIn/google/domain"
)

type getAccessTokenRepository struct {
	GoogleConfig domain.GoogleConfig
}

func New(gc domain.GoogleConfig) domain.GetAccessTokenRepository {
	return &getAccessTokenRepository{GoogleConfig: gc}
}

func (g *getAccessTokenRepository) GetAccessToken(code string) (domain.GetAccessTokenResponse, error) {
	rtn := domain.GetAccessTokenResponse{}

	requestBody := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {g.GoogleConfig.RedirectUrl},
		"client_id":     {g.GoogleConfig.ClientId},
		"client_secret": {g.GoogleConfig.ClientSecret},
	}

	response, err := http.PostForm(g.GoogleConfig.TokenApi, requestBody)
	if err != nil {
		return rtn, err
	}

	defer response.Body.Close()

	var responseDecode domain.GetAccessTokenResponse
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
