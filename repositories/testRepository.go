package repositories

import (
	"go-api/model"
	"go-api/servers"
)

type TestRepository struct {

}

var TestRepo TestRepository


func (self TestRepository) GetTest()(rsltArr[] model.Member){
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

func (self TestRepository) InsertMember(member[] model.Member)(){
	dbmap:= servers.Database().GetInstance()
	dbmap.Insert()
}
