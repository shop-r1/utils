package models

import "time"

type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	Scopes       []string  `json:"scopes"`
}

type AuthorizeRequest struct {
	ResponseType string   `json:"response_type"`
	ClientId     uint     `json:"client_id"`
	Scopes       []string `json:"scopes"`
	State        string   `json:"state"`
	RedirectUri  string   `json:"redirect_uri"`
}

type AuthorizeResponse struct {
	Code  string `json:"code"`
	State string `json:"state"`
	Token *Token `json:"token"`
}

type TokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	// password, client_credentials, authorization_code, refresh_token
	GrantType    string `json:"grant_type"`
	RedirectUri  string `json:"redirect_uri"`
	RefreshToken string `json:"refresh_token,omitempty"`
	// scopes can be added for client_credentials request
	Scopes []string `json:"scopes"`
	// metadata to be stored with a token that's generated
	Metadata interface{} `json:"metadata"`
}

type TokenResponse struct {
	Token *Token `json:"token"`
}

type RevokeRequest struct {
	// revoke access token
	AccessToken string `json:"access_token"`
	// revoke via refresh token
	RefreshToken string `json:"refresh_token"`
}

type IntrospectRequest struct {
	TenantId    uint   `json:"tenant_id"`
	AccessToken string `json:"access_token"`
}

type IntrospectResponse struct {
	Related *Related `json:"related"`
	Active  bool     `json:"active"`
}

type Related struct {
	TenantId uint     `json:"tenant_id"`
	RoleId   uint     `json:"role_id"`
	UserId   uint     `json:"user_id"`
	Scopes   []string `json:"scopes"`
}
