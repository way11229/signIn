package getAccessTokenRepository

import (
	"encoding/json"
	"errors"
	"net/http"
	"signIn/fb/domain"
)

type getAccessTokenRepository struct {
	FbConfig domain.FbConfig
}

func New(fc domain.FbConfig) domain.GetAccessTokenRepository {
	return &getAccessTokenRepository{FbConfig: fc}
}

func (g *getAccessTokenRepository) GetAccessToken(code string) (domain.GetAccessTokenResponse, error) {
	rtn := domain.GetAccessTokenResponse{}

	requestUrl := g.FbConfig.TokenApi + "?client_id=" + g.FbConfig.ClientId + "&redirect_uri=" + g.FbConfig.RedirectUrl + "&client_secret=" + g.FbConfig.ClientSecret + "&code=" + code
	request, requstErr := http.NewRequest("GET", requestUrl, nil)
	if requstErr != nil {
		return rtn, requstErr
	}

	client := &http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return rtn, responseErr
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var errorDecode domain.ErrorResponse
		decodeErr := json.NewDecoder(response.Body).Decode(&errorDecode)
		if decodeErr != nil {
			return rtn, decodeErr
		}

		return rtn, errors.New(errorDecode.Error.Message)
	}

	var responseDecode domain.GetAccessTokenResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseDecode)
	if decodeErr != nil {
		return rtn, decodeErr
	}

	rtn = responseDecode

	return rtn, nil
}
