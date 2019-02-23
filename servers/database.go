package servers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type dbInstance  struct {
}

var TestDB dbInstance


func (self dbInstance) InitDB() *gorp.DbMap{
	// db, dbmap 객체 생성
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/temp")
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}
