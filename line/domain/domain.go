package domain

type LineConfig struct {
	TokenApi     string `json:"tokenApi"`
	VerifyApi    string `json:"verifyApi"`
	ProfileApi   string `json:"profileApi"`
	RedirectUrl  string `json:"redirectUrl"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpireIn         uint32 `json:"expire_in"`
	IdToken          string `json:"id_token"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type VerifyIdTokenResponse struct {
	Iss              string   `json:"iss"`
	Sub              string   `json:"sub"`
	Aud              string   `json:"aud"`
	Exp              uint32   `json:"exp"`
	Iat              uint32   `json:"iat"`
	Nonce            string   `json:"nonce"`
	Amr              []string `json:"amr"`
	Name             string   `json:"name"`
	Picture          string   `json:"picture"`
	Email            string   `json:"email"`
	Error            string   `json:"error"`
	ErrorDescription string   `json:"error_description"`
}

type UserProfileResponse struct {
	UserId           string `json:"userId"`
	DisplayName      string `json:"displayName"`
	PictureUrl       string `json:"pictureUrl"`
	StatusMessage    string `json:"statusMessage"`
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
	StatusMessage       string `json:"statusMessage"`
}

type ServiceList struct {
	SignInService SignInService
}

type SignInServiceRepositoryList struct {
	GetAccessTokenRepository GetAccessTokenRepository
	VerifyIdTokenRepository  VerifyIdTokenRepository
	GetUserProfileRepository GetUserProfileRepository
}

type SignInService interface {
	SignIn(string) (SignInResponse, error)
}

type GetAccessTokenRepository interface {
	GetAccessToken(string) (AccessTokenResponse, error)
}

type VerifyIdTokenRepository interface {
	VerifyIdToken(string) (VerifyIdTokenResponse, error)
}

type GetUserProfileRepository interface {
	GetUserProfile(string) (UserProfileResponse, error)
}
