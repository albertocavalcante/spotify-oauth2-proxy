package oauth2

type OAuthData struct {
	RefreshToken string
	BearerToken  string
	OauthExpires string
	OauthHMAC    string
}

type ServerConfig struct {
	Port string
	FrontendStaticFiles string
}
