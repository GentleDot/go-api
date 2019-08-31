package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
	"log"
	"time"
)


func OauthRoute(){
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	oauthServer := router.Group("/oauth")
	{
		controller := controllers.OauthServer

		oauthServer.GET("/login", controller.Login)
		oauthServer.GET("/auth", controller.Auth )
		oauthServer.GET("/authorize", controller.Authorize )
		oauthServer.GET("/token", controller.Token )
		oauthServer.GET("/test", controller.Test )

	}

	log.Println("Client is running at 9096 port.")

	router.Run(":9096")
}