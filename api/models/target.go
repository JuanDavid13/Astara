package models

import (
	//"fmt"
	//"encoding/json"
	//"strconv"
	//C "astara/commons"
	//re "reflect"
	"database/sql"
)

type Target struct {
	id int`json:id`;
	Name string `json:name`;
	//Unit string `json:unit`;
	//Quantity int`json:quantity`;
	Priority int`json:priority`;
	//Id_elem int`json:id_elem`;
}

func (t *Target) GetTargets(db *sql.DB,uid int) {
	_ , err := db.Prepare("SELECT * FROM `Targets` WHERE Id LIKE ?;");
	if err != nil{panic(err);}
}

func (t *Target) GetAllTargets(db *sql.DB) []Target {
	rows, err := db.Query("SELECT * FROM `Targets`;");
	if err != nil{panic(err);}


	//var (
	//	id sql.NullInt64;
	//	name sql.NullString;
	//	priority sql.NullInt32;
	//)

	type column struct {
		id				sql.NullInt64;
		name			sql.NullString;
		priority	sql.NullInt32;
	}	
	col := column{}
	target := Target{}
	//var results	[]string;
	results	:= []Target{};

	for rows.Next(){
		rows.Scan(&col.id,&col.name,&col.priority);
		//fmt.Printf("%d - %s (%d)\n",col.id.Int64,col.name.String,col.priority.Int32);
		target.id = int(col.id.Int64);
		target.Name = col.name.String;
		target.Priority = int(col.priority.Int32);

		results = append(results, target);

		//base 10 = decimal
		//results = append(results,strconv.FormatInt(col.id.Int64,10),col.name.String, strconv.FormatInt(int64(col.priority.Int32),10))
	}

	//result, _ := json.Marshal(results);
	//fmt.Println(string(result));

	return (results)

}

