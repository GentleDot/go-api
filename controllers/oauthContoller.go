package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"time"
)

type oauthContoller struct {

}

const (
	authServerURL = "http://localhost:9096"
)

var (
	Oauth oauthContoller
	config = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "22222222",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9094/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/authorize",
			TokenURL: authServerURL + "/token",
		},
	}
	globalToken *oauth2.Token
)
func (self oauthContoller) Init(c *gin.Context) {
	u := config.AuthCodeURL("xyz")
	c.Redirect(http.StatusFound, u)

}

func (self oauthContoller) Oauth2(c *gin.Context){
	r := c.Request
	r.ParseForm()

	state := r.Form.Get("state")

	if state != "xyz" {
		c.String(http.StatusBadRequest, "State invalid")
		//http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}
	code := r.Form.Get("code")
	if code == "" {

		c.String(http.StatusBadRequest, "Code not found")
		//http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error());
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	globalToken = token

	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "  ")
	e.Encode(token)

}

func (self oauthContoller) Refresh(c *gin.Context){
	if globalToken == nil {
		c.Redirect(http.StatusFound, "/")
		//http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	globalToken.Expiry = time.Now()
	token, err := config.TokenSource(context.Background(), globalToken).Token()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error());
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalToken = token
	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "  ")
	e.Encode(token)
}

func (self oauthContoller) Try(c *gin.Context){
	if globalToken == nil {
		c.Redirect(http.StatusFound, "/")
		//http.Redirect(w, r, "/", http.StatusFound)
		return
	}


	resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	io.Copy(c.Writer, resp.Body)
}

func (self oauthContoller) Pwd(c *gin.Context) {
	token, err := config.PasswordCredentialsToken(context.Background(), "test", "test")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error());
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalToken = token
	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "  ")
	e.Encode(token)
}

func (self oauthContoller) Client(c *gin.Context) {
	cfg := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.Endpoint.TokenURL,
	}

	token, err := cfg.Token(context.Background())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error());
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(c.Writer)
	e.SetIndent("", "  ")
	e.Encode(token)
}
