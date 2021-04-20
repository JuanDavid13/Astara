package commons;

import (
	"database/sql"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct{
	driver, dbName, user, pwd string;
}


func (Db *Db) Open(user, pwd string) *sql.DB {
	Db.driver = "mysql";
	Db.dbName = "astara";

	db, err := sql.Open(Db.driver,user+":"+pwd+"@/"+Db.dbName);
	if err != nil{panic(err)};
	return db;
}
