package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
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

	router.Run()
}