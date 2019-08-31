package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api/apiConfig"
	"go-api/controllers"
	"log"
	"time"
)


func Route()  {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))



	member := router.Group(apiConfig.PREFIX_URI + "member")
	{
		member.GET("/test", controllers.Test.FuncReqTest)
		member.POST("/newMember", controllers.Member.InsMemberData)
	}

	oauth := router.Group(apiConfig.PREFIX_URI + "oauthClient")
	{
		oauthClientController := controllers.Oauth

		oauth.GET("/", oauthClientController.Init)
		oauth.GET("/oauth2", oauthClientController.Oauth2)
		oauth.GET("/refresh", oauthClientController.Refresh)
		oauth.GET("/try", oauthClientController.Try)
		oauth.GET("/pwo", oauthClientController.Pwd)
		oauth.GET("/client", oauthClientController.Client)

		//log.Fatal(http.ListenAndServe(":8080", nil))
		log.Println("클라이언트 라우터 동작!")
	}

	oauthServer := router.Group(apiConfig.PREFIX_URI + "oauthServer")
	{
		oauthServerController := controllers.OauthServer

		oauthServer.GET("/login", oauthServerController.Login)
		oauthServer.GET("/auth", oauthServerController.Auth )
		oauthServer.GET("/authorize", oauthServerController.Authorize )
		oauthServer.GET("/token", oauthServerController.Token )
		oauthServer.GET("/test", oauthServerController.Test )

		log.Println("서버 라우터 동작!")
	}

	router.Run()
}