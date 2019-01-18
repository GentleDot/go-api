package controllers

import "github.com/gin-gonic/gin"

func FuncReq(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message" : "test!",
	})
}

func FuncReqTest(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message" : "/api/v1/member/test 주소를 호출한 결과입니다.",
	})
}