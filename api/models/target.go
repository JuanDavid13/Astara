package models

import (
	"fmt"
	//C "astara/commons"
	//re "reflect"
	"database/sql"
)

type Target struct {
	Name, Unit string;
	Quantity, Priority, Id, Id_elem int;
}

func (t *Target) GetTargets(db *sql.DB,uid int) {
	_ , err := db.Prepare("SELECT * FROM `Targets` WHERE Id LIKE ?;");
	if err != nil{panic(err);}
}

func (t *Target) GetAllTargets(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM `Targets`;");
	if err != nil{panic(err);}

	var (
		id sql.NullInt64;
		name sql.NullString;
		priority sql.NullInt32;
	)
	
	for rows.Next(){
		rows.Scan(&id,&name,&priority)
		fmt.Printf("%d - %s (%d)\n",id.Int64,name.String,priority.Int32);
	}
}

