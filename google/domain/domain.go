package domain

type GoogleConfig struct {
	TokenApi     string `json:"tokenApi"`
	UserInfoApi  string `json:"userInfoApi"`
	RedirectUrl  string `json:"redirectUrl"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        uint32 `json:"expires_in"`
	IdToken          string `json:"id_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type GetUserInfoResponse struct {
	Sub              string `json:"sub"`
	Name             string `json:"name"`
	GivenName        string `json:"given_name"`
	FamilyName       string `json:"family_name"`
	Picture          string `json:"picture"`
	Email            string `json:"email"`
	EmailVerified    bool   `json:"email_verified"`
	Locale           string `json:"locale"`
	Hd               string `json:"hd"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
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
	GetUserInfoRepository    GetUserInfoRepository
}

type SignInService interface {
	SignIn(string) (SignInResponse, error)
}

type GetAccessTokenRepository interface {
	GetAccessToken(string) (GetAccessTokenResponse, error)
}

type GetUserInfoRepository interface {
	GetUserInfo(string) (GetUserInfoResponse, error)
}
