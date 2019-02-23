package repositories

import (
	"go-api/servers"
	"log"
)

type Test struct {
	TESTKEY string
	TESTVALUE string
}

type testRepository struct {

}

var TestRepository testRepository


func (self testRepository) GetTest(pkStr string)(rsltArr []string){
	var rslt Test
	// db 접속
	dbmap := servers.TestDB.InitDB()
	dbmap.AddTableWithName(Test{}, "TEST")
	defer dbmap.Db.Close()

	// 쿼리
	strQuery := "select TESTKEY, TESTVALUE from TEST where TESTKEY=?"

	err := dbmap.SelectOne(&rslt, strQuery, pkStr)
	checkErr(err, "테이블 조회 실패!")

	rsltArr = append(rsltArr, rslt.TESTKEY)
	rsltArr = append(rsltArr, rslt.TESTVALUE)

	return rsltArr
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
