package services

type responseStr  struct {
}

var TestService responseStr

func (self responseStr) TestSer()(msg string) {
	msg = "testStr"
}