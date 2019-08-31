package services

import (
	"go-api/apiConfig"
	"golang.org/x/oauth2"
)


type OauthClientService struct {

}

const (
	AUTH_SERVER_URL = "http://localhost:8080" + apiConfig.PREFIX_URI + "oauthServer"
)

var (
	OauthClient OauthClientService
	OauthConfig = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "22222222",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:8080" + apiConfig.PREFIX_URI + "oauthClient",
		Endpoint: oauth2.Endpoint{
			AuthURL:  AUTH_SERVER_URL + "/authorize",
			TokenURL: AUTH_SERVER_URL + "/token",
		},
	}
)



