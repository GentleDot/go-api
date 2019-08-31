package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"go-api/services"
	"log"
	"net/http"
	"net/url"
	"time"
)

type OauthServerController struct {

}

var (
	OauthServer OauthServerController
	oauthServerService = services.OauthServerService
	srv = oauthServerService.Main()
)
func (self OauthServerController) Login (c *gin.Context) {
	w := c.Writer
	r := c.Request
	store, err := session.Start(nil, w, r)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if c.Request.Method == "POST" {
		store.Set("LoggedInUserID", "000000")
		store.Save()

		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}

	//outputHTML(w, r, "static/login.html")
	c.Redirect(http.StatusPermanentRedirect, "static/login.html")

}


func (self OauthServerController) Auth (c *gin.Context){
	w := c.Writer
	r := c.Request
	store, err := session.Start(nil, w, r)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok{

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)

		return
	}

	//outputHTML(w, r, "static/auth.html")
	c.Redirect(http.StatusPermanentRedirect, "static/auth.html")

}

func (self OauthServerController) Authorize (c *gin.Context){
	w := c.Writer
	r := c.Request
	store, err := session.Start(nil, w, r)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("store : ")
	log.Println(store)
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values);
	}

	r.Form = form

	log.Println("Form : ")
	log.Println(r.Form)

	log.Println("store : ")
	log.Println(r.Form)

	store.Delete("ReturnUri")
	store.Save()

	err = srv.HandleAuthorizeRequest(w, r)

	log.Println("request : ")
	log.Println(r)
	log.Println("writer : ")
	log.Println(w)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
}

func (self OauthServerController) Token (c *gin.Context) {
	w := c.Writer
	r := c.Request

	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (self OauthServerController) Test (c *gin.Context) {
	w := c.Writer
	r := c.Request

	token, err := srv.ValidationBearerToken(r)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}
