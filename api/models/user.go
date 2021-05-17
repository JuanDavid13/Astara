package models

import (
	"fmt"
	//"reflect"
	//"os"

	db "astara/commons/database"
	"database/sql"
	//jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
  Id int `json:"id"`;
  User string `json:"user"`;
  Pass string `json:"pwd"`;
}

func IsRegistered(user string) bool {
  db := db.GetDb("nonuser");
  query := "SELECT `id`  FROM `Users` WHERE (`Name` LIKE ? OR `Email` LIKE ?)";
  stmt, err := db.Prepare(query);
  if err != nil{
    panic(err);
  }

  var id int;
  err = stmt.QueryRow(user, user).Scan(&id);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows {
    return false;
    //panic(err);
  }
  if id != 0 {return true;}
  return false;
}

func CheckCredentials(user, pass string) (int,string){
  fmt.Println("is registered:")

  db := db.GetDb("nonuser");
  query := "SELECT U.`Id`, R.`Name` FROM `Users` AS U JOIN `Rols` AS R ON (U.`Id_rol` = R.`Id` ) WHERE (U.`Name` LIKE ? OR U.`Email` LIKE ?) AND U.`Password` LIKE ?";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  var ( 
    id int;
    rol string; 
  )
  err = stmt.QueryRow(user,user,pass).Scan(&id,&rol);
  defer stmt.Close();
  //if err != nil { panic(err); }
  if err != nil && err != sql.ErrNoRows {
    return -1,"";
    //panic(err);
  }

  if id != 0 { return id,rol; }
  return -1,"";
}
