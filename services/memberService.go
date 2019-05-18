package services

import (
	"go-api/model"
	"go-api/repositories"
)

type memberService struct {

}

var MemberServ memberService
var memberRep repositories.MemberRepository

func (self memberService) InsertMember(memNm, memId string)(rslt error) {
	memberData := model.NewMember(memNm, memId)
	rslt = memberRep.InsertMember(memberData)

	return rslt
}