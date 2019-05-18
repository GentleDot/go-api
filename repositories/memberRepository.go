package repositories

import (
	"go-api/model"
	"go-api/servers"
)

type MemberRepository struct {

}

var MemberRep MemberRepository


func (self MemberRepository) GetTest()(rsltArr[] model.Member){
	// db 접속
	dbmap := servers.Database().GetInstance()
	dbmap.AddTableWithName(model.Member{}, "Member").SetKeys(true, "member_no")

	// 쿼리 selectOne
	//strQuery := "select * from member where TESTKEY=?"
	strQuery := "select * from member"

	// selectOne
	//err := dbmap.SelectOne(&rslt, strQuery, pkStr)
	_, err := dbmap.Select(&rsltArr, strQuery)
	servers.CheckErr(err, "테이블 조회 실패!")


	return rsltArr
}


func (self *MemberRepository) InsertMember(member model.InsMember) error {

	dbmap:= servers.Database().GetInstance()
	dbmap.AddTableWithName(model.InsMember{}, "Member").SetKeys(true, "member_no")
	trans, err := dbmap.Begin()

	if err != nil{
		return err
	}

	err = trans.Insert(&member)
	servers.CheckErr(err, "저장에 실패하였습니다.")

	return trans.Commit()

}


