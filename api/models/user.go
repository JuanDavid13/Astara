package models

import (
	//"os"

	db "astara/commons/database"
	//jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
  Id int `json:"id"`;
  User string `json:"user"`;
  Pass string `json:"pwd"`;
}

func IsRegistered(user string) bool {
  //Db := db.Db{};
  //Db.New();
  //db := Db.Open(os.Getenv("DB_NOUSER_USER"), os.Getenv("DB_NOUSER_PWD"));
  db := db.GetDb();
  //defer db.Close();

  query := "SELECT id FROM `Users` WHERE (`Name` LIKE ? OR `Email` LIKE ?)";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  row := stmt.QueryRow(user, user);
  defer stmt.Close();

  var id int;
  row.Scan(&id)

  if(id != 0){return true;}

  return false;
}

func CheckCredentials(user, pass string) int {
  //Db := db.Db{};
  //Db.New();
  //db := Db.Open(os.Getenv("DB_NOUSER_USER"), os.Getenv("DB_NOUSER_PWD"));
  db := db.GetDb();
  //defer db.Close();

  query := "SELECT id FROM `Users` WHERE (`Name` LIKE ? OR `Email` LIKE ?) AND `Password` LIKE ?";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  row := stmt.QueryRow(user,user,pass);
  defer stmt.Close();

  var id int;
  row.Scan(&id)

  if(id != 0){return id;}

  return -1;
}
