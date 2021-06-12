package domain

type LineConfig struct {
	TokenApi     string `json:"tokenApi"`
	VerifyApi    string `json:"verifyApi"`
	RedirectUrl  string `json:"redirectUrl"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpireIn     uint32 `json:"expire_in"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type UserDataResponse struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     uint32   `json:"exp"`
	Iat     uint32   `json:"iat"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
}

type SignInResponse struct {
	AccessToken         string `json:"accessToken"`
	AccessTokenExpireIn uint32 `json:"accessTokenExpireIn"`
	RefreshToken        string `json:"refreshToken"`
	UserId              string `json:"userId"`
	Name                string `json:"name"`
	Picture             string `json:"picture"`
	Email               string `json:"email"`
}

type ServiceList struct {
	SignInService SignInService
}

type SignInServiceRepositoryList struct {
	GetAccessTokenRepository GetAccessTokenRepository
	GetUserDataRepository    GetUserDataRepository
}

type SignInService interface {
	SignIn(string) (SignInResponse, error)
}

type GetAccessTokenRepository interface {
	GetAccessToken(string) (AccessTokenResponse, error)
}

type GetUserDataRepository interface {
	GetUserData(string) (UserDataResponse, error)
}
