package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/model"
	"go-api/services"
)

type TestController struct {

}

var Test TestController



func (self TestController) FuncReqTest(c *gin.Context)  {
	var members[] model.Member

	members = services.TestService.TestSer()

	c.JSON(200, gin.H{
		"members" : members,
	})
}