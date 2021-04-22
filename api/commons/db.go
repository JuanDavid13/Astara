package commons;

import (
	"database/sql"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct{
	driver, dbName, user, pwd string;
}

//"Constructor"
func (Db *Db) New(){
	Db.driver = "mysql";
	Db.dbName = "astara";
}


func (Db *Db) Open(user, pwd string) *sql.DB {
	db, err := sql.Open(Db.driver,user+":"+pwd+"@/"+Db.dbName);
	if err != nil{panic(err)};
	return db;
}

func (Db *Db) CheckPool(db *sql.DB) bool {
	if db.Ping() == nil{
		return true;
	}else{
		return false;
	}
}

