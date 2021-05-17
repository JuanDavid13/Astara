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
	NonUserInstance *sql.DB; //logger
	UserInstance *sql.DB;
	AdminUserInstance *sql.DB;
	//loggerInstance *sql.DB;
)

var lock = &sync.Mutex{};
//var once sync.Once;

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

	//instance := db.open(user,pwd);
	//setInstance(instance, group);

	//return instance;
	return db.open(user,pwd);
}


func (db *db) open(user, pwd string) *sql.DB {
	if db, err := sql.Open(db.driver,user+":"+pwd+"@/"+db.dbName); err != nil {
		panic(err);
	}else{ 
		fmt.Println("lo que devuelve open:");
		fmt.Println(db);
		return db
	}
}

//borrar
func getInstance(rol string) *sql.DB{
	switch rol {
		case "Nonuser": { return NonUserInstance; }
		case "user": { return UserInstance; }
		case "admin": { return AdminUserInstance; }
		default: { return nil; }
	}
}

func GetDb(rol string) *sql.DB {
	//fmt.Println("NonUserInstance");
	//fmt.Println(NonUserInstance);
	//fmt.Println("UserInstance");
	//fmt.Println(UserInstance);
	//fmt.Println("AdminUserInstance");
	//fmt.Println(AdminUserInstance);

	switch rol {
		case "nonuser": { 
			if NonUserInstance == nil {
				//once.Do(func(){ 
				//	newInstance := *db{}.newdb(rol)
				//	NonUserInstance = &newInstance; });
				lock.Lock();
				defer lock.Unlock();
				if NonUserInstance == nil	{
					//fmt.Println("creating NonUserInstance");
					newInstance := *db{}.newdb(rol);
					NonUserInstance = &newInstance;
				}
			}//else{ fmt.Println("Already has been created an NonUserInstance"); }
			return NonUserInstance;
		}
		case "user": { 
			if UserInstance == nil {
				//once.Do(func(){ 
				//	newInstance := *db{}.newdb(rol)
				//	UserInstance = &newInstance; });
				lock.Lock();
				defer lock.Unlock();
				if UserInstance == nil	{
					//fmt.Println("creating UserInstance");
					newInstance := *db{}.newdb(rol);
					UserInstance = &newInstance;
				}
			}//else{ fmt.Println("Already has been created an UserInstance"); }
			return UserInstance;
		}
		case "admin": {
			if AdminUserInstance == nil {
				//once.Do(func(){
				//	newInstance := *db{}.newdb(rol)
				//	AdminUserInstance = &newInstance; });
				lock.Lock();
				defer lock.Unlock();
				if AdminUserInstance == nil	{
					//fmt.Println("creating AdminUserInstance");
					newInstance := *db{}.newdb(rol);
					AdminUserInstance = &newInstance;
				}
			}//else{ fmt.Println("Already has been created an AdminUserInstance"); }
			return AdminUserInstance;
		}
		default: { return nil; }
	}
}
