package services

import "go-api/repositories"

type responseStr  struct {

}

var TestService responseStr


func (self responseStr) TestSer()(msg string) {
	var testVal []string
	testVal = repositories.TestRepository.GetTest("TEST1")

	msg = "테스트 결과 : 키는 " + testVal[0] + " 값은 " + testVal[1] + "입니다."
	return msg
}