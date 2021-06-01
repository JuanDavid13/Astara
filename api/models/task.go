package models;

import (
	"fmt"
	//"reflect"
	//"os"

	db "astara/commons/database"
	"database/sql"
)

func CreateNewTask(uid, areaid int, rol, name, deadline, dated string) bool {
  fmt.Println("Create Task:");
  db := db.GetDb(rol);
  
	targetQuery := "INSERT INTO `Targets` (`Id_usu`,`Id_area`,`Name`,`Deadline`) VALUES(?, ?, ?, ?);";

  targetStmt, err := db.Prepare(targetQuery);
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err); }

	res, err := targetStmt.Exec(uid, areaid, name, deadline);
  defer targetStmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err);}

	targetId, err := res.LastInsertId();
	if err != nil { return false; }

	taskQuery := "INSERT INTO `Task` (`Id_target`,`Dated`) VALUES(?, ?);";

  taskStmt, err := db.Prepare(taskQuery);
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err); }

	_ , err = taskStmt.Exec(targetId, dated);
  defer taskStmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err);}

	return true;
}
