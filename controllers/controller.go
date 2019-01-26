package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/services"
)

type TestController struct {

}


var Test TestController

func (self TestController) FuncReqTest(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message" : services.TestService.TestSer(),
	})
}