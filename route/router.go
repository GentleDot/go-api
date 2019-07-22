package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-glog"
	"github.com/zalando/gin-oauth2"
	"github.com/zalando/gin-oauth2/zalando"
	"go-api/controllers"
	"go-api/model"
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

	aouth := router.Group("api/v1/aouth")
	{
		aouth.Use(ginglog.Logger(3 * time.Second))
		aouth.Use(ginoauth2.RequestLogger([]string{"uid"}, "data"))
		aouth.Use(gin.Recovery())

		aouth.Use(ginoauth2.Auth(zalando.UidCheck(model.USERS), zalando.OAuth2Endpoint))

		aouth.GET("/", controllers.AouthContoller.ReqToken)
	}

	router.Run()
}