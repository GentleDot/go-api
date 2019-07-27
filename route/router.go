package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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



	member := router.Group("/api/v1/member")
	{
		member.GET("/test", controllers.Test.FuncReqTest)
		member.POST("/newMember", controllers.Member.InsMemberData)
	}

	oauth := router.Group("api/v1/oauth")
	{

		oauth.GET("/", controllers.Oauth.Init)
		oauth.GET("/oauth2", controllers.Oauth.Oauth2)
		oauth.GET("/refresh", controllers.Oauth.Refresh)
		oauth.GET("/try", controllers.Oauth.Try)
		oauth.GET("/pwo", controllers.Oauth.Pwd)
		oauth.GET("/client", controllers.Oauth.Client)

		log.Println("Client is running at 8080 port.")
		//log.Fatal(http.ListenAndServe(":8080", nil))
	}

	router.Run()
}