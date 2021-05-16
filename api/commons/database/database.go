package database;

import (
	"database/sql"
	"os"

	"sync"
	"fmt"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)


type db struct{
	driver, dbName, user, pwd string;
}

var (
	nonUserInstance *sql.DB;
	userInstance *sql.DB;
	adminUserInstance *sql.DB;
	loggerInstance *sql.DB;
)

var once sync.Once;

//"Constructor"
func (db db) newdb(group string) *sql.DB{
	db.driver = "mysql";
	db.dbName = os.Getenv("DB_NAME");

	var (
		user, pwd string
	)

	switch group {
		case "nouser": { 
			user = os.Getenv("DB_NOUSER_USER");
			pwd = os.Getenv("DB_NOUSER_PWD");
		}
		case "user": {
			user = os.Getenv("DB_USER_USER");
			pwd = os.Getenv("DB_USER_PWD");
		}
		case "admin": {
			user = os.Getenv("DB_ADMIN_USER");
			pwd = os.Getenv("DB_ADMIN_PWD");
		}
		default: {
			user = os.Getenv("DB_NOUSER_USER");
			pwd = os.Getenv("DB_NOUSER_PWD");
		}
	}//breaks are added automatically
		
	return db.open(user,pwd);
}


func (db *db) open(user, pwd string) *sql.DB {
	if db, err := sql.Open(db.driver,user+":"+pwd+"@/"+db.dbName); err != nil {
		panic(err);
	}else{ return db }
}

func (Db *db) CheckPool(db *sql.DB) bool {
	if db.Ping() == nil{
		return true;
	}else{
		return false;
	}
}

func GetDb() *sql.DB {
	if userInstance == nil {
		once.Do(func(){
			fmt.Println("Creating a instance");
			DB := db{}.newdb("user");
			userInstance = DB;
		});
	}else{
		fmt.Println("Already has been created");
	}
	return userInstance
}
