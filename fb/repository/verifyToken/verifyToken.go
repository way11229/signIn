package verifyToken

import (
	"encoding/json"
	"errors"
	"net/http"
	"signIn/fb/domain"
)

type verifyTokenRepository struct {
	FbConfig domain.FbConfig
}

func New(fc domain.FbConfig) domain.VerifyTokenRepository {
	return &verifyTokenRepository{FbConfig: fc}
}

func (g *verifyTokenRepository) VerifyToken(accessToken string) (domain.VerifyTokenResponse, error) {
	rtn := domain.VerifyTokenResponse{}

	requestUrl := g.FbConfig.VerifyApi + "?input_token=" + accessToken + "&access_token=" + accessToken
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

	var responseDecode domain.VerifyTokenResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseDecode)
	if decodeErr != nil {
		return rtn, decodeErr
	}

	rtn = responseDecode

	return rtn, nil
}
