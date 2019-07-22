package model

import (
	"github.com/zalando/gin-oauth2/zalando"
)

type aouthMember struct {

}

var USERS []zalando.AccessTuple = []zalando.AccessTuple{
	{"/employees", "sszuecs", "Sandor Szücs"},
	{"/employees", "njuettner", "Nick Jüttner"},
}