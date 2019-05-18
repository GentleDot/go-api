package services

import (
	"go-api/model"
	"go-api/repositories"
)

type responseStr  struct {

}

var TestService responseStr
var repository repositories.TestRepository


func (self responseStr) TestSer()(testVal[] model.Member) {
	testVal = repository.GetTest()

	return testVal
}