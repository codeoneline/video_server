package dbops

import "database/sql"

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn, err = sql.Open("mysql", "root:jkl456@tcp(localhost:3306)/mygo?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
