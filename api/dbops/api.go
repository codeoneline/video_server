package dbops

import (
	"database/sql"
	_ "github.com/go-driver/mysql"
)

// func openConn() *sql.DB {
// 	dbConn, err := sql.Open("mysql", "root:jkl456@tcp(localhost:3306)/mygo?charset=utf8")
// }

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}