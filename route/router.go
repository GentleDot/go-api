package route

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

func Route()  {
	router := gin.Default()

	member := router.Group("/api/v1/member")
	{
		member.GET("/", controllers.FuncRes)
	}

	router.Run()
}