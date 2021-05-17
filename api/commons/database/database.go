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
	nonUserInstance *sql.DB; //logger
	userInstance *sql.DB;
	adminUserInstance *sql.DB;
	//loggerInstance *sql.DB;
)

var once sync.Once;

//"Constructor"
func (db db) newdb(group string) *sql.DB{

	db.driver = "mysql";
	db.dbName = os.Getenv("DB_NAME");

	var (	user, pwd string )

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

	instance := db.open(user,pwd);
	//setInstance(instance, group);

	return instance;
}


func (db *db) open(user, pwd string) *sql.DB {
	if db, err := sql.Open(db.driver,user+":"+pwd+"@/"+db.dbName); err != nil {
		panic(err);
	}else{ return db }
}

func setInstance(instance *sql.DB, rol string) {
	fmt.Println("set instance");
	switch rol {
		case "nouser": { nonUserInstance = instance	}
		case "user": { userInstance = instance	}
		case "admin": { adminUserInstance = instance }
		default: {}
	}

	fmt.Println(&instance);
}

func GetDb(rol string) *sql.DB {
	fmt.Println("get db instance");

	var instance **sql.DB;

	switch rol {
		case "nonuser":{ instance = &nonUserInstance; } //logger
		case "user":{ instance = &userInstance; }
		case "admin":{ instance = &adminUserInstance; }
		default: { instance = &nonUserInstance; }
	}

	if *instance == nil {
		once.Do(func(){
			fmt.Println("Creating a instance of type", rol);
			newInstance := db{}.newdb(rol)
			//instance = db{}.newdb(rol);
			*instance = newInstance;
		});
	}else{ 
		fmt.Println("Already has been created"); 
	}

	fmt.Println("instance before return");
	fmt.Println(instance);

	return *instance;
}
