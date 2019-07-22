package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type aouthContoller struct {

}

var AouthContoller aouthContoller

func (self aouthContoller) ReqToken(c *gin.Context) {

	fmt.Print(c.Errors.JSON());
	fmt.Println(c.ContentType());
	if v, ok := c.Get("cn"); ok {
		c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for users to %s", v)})
	} else {
		c.JSON(200, gin.H{"message": "Hello from private for users without cn"})
	}


}
