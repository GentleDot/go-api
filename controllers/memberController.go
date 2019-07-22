package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/model"
	"go-api/services"
)

type MemberController struct {

}

var Member MemberController



func (self MemberController) FuncReqMemberData(c *gin.Context)  {
	var members[] model.Member

	members = services.TestService.TestSer()

	c.JSON(200, gin.H{
		"members" : members,
	})
}

func (self MemberController) InsMemberData(c *gin.Context)  {
	var member model.InsMember

	c.ShouldBindJSON(&member)

	memberNm := member.MemberName
	memberId := member.MemberId

	fmt.Println("id :" + memberId + " Nm :" + memberNm)

	rslt := services.MemberServ.InsertMember(memberNm, memberId)

	c.JSON(200, gin.H{
		"result" : rslt,
	})
}