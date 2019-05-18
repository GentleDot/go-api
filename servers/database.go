package servers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type database  struct {
	dbMap *gorp.DbMap
}

var instance *database
var once sync.Once

func Database() *database{
	once.Do(func() {
		// db, dbmap 객체 생성
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
		CheckErr(err, "Database Open failed.")
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		instance = &database{dbMap : dbmap}
	})

	return instance
}

func (db *database) GetInstance() *gorp.DbMap {
	return db.dbMap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
