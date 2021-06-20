package domain

type FbConfig struct {
	TokenApi     string `json:"tokenApi"`
	VerifyApi    string `json:"verifyApi"`
	ProfileApi   string `json:"profileApi"`
	RedirectUrl  string `json:"redirectUrl"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type ErrorResponse struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      uint32 `json:"code"`
		FbTraceId string `json:"fbtrace_id"`
	} `json:"error"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint32 `json:"expires_in"`
}

type VerifyTokenResponse struct {
	Data struct {
		AppId       uint64 `json:"app_id"`
		Type        string `json:"type"`
		Application string `json:"application"`
		ExpiresAt   uint32 `json:"expires_at"`
		IsValid     bool   `json:"isValid"`
		IssuedAt    uint32 `json:"issue_at"`
		Metadata    struct {
			Sso string `json:"sso"`
		} `json:"metadata"`
		Scopes []string `json:"scopes"`
		UserId string   `json:"user_id"`
	} `json:"data"`
}

type PictureContent struct {
	Data struct {
		Height        uint32 `json:"height"`
		Is_silhouette bool   `json:"is_silhouette"`
		Url           string `json:"url"`
		Width         uint32 `json:"width"`
	} `json:"data"`
}

type GetUserProfileResponse struct {
	UserId   string         `json:"id"`
	Name     string         `json:"name"`
	Picture  PictureContent `json:"picture"`
	Email    string         `json:"email"`
	Birthday string         `json:"birthday"`
}

type SignInResponse struct {
	AccessToken         string `json:"accessToken"`
	AccessTokenExpireIn uint32 `json:"accessTokenExpireIn"`
	UserId              string `json:"userId"`
	Name                string `json:"name"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
	Birthday            string `json:"birthday"`
}

type ServiceList struct {
	SignInService SignInService
}

type SignInServiceRepositoryList struct {
	GetAccessTokenRepository GetAccessTokenRepository
	VerifyTokenRepository    VerifyTokenRepository
	GetUserProfileRepository GetUserProfileRepository
}

type SignInService interface {
	SignIn(string) (SignInResponse, error)
}

type GetAccessTokenRepository interface {
	GetAccessToken(string) (GetAccessTokenResponse, error)
}

type VerifyTokenRepository interface {
	VerifyToken(string, string) (VerifyTokenResponse, error)
}
type GetUserProfileRepository interface {
	GetUserProfile(string, string) (GetUserProfileResponse, error)
}
