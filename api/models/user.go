package models

import (
	"os"

	db "astara/database"

	//jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
  Id int `json:"id"`;
  Name string `json:"name"`;
  Pass string `json:"pwd"`;
  //Token jwt.Token `json:token`;
}

//func Authorice (name, pwd string) *jwt.Token {
    
//}

func (u *User) IsRegistered() bool {
  Db := db.Db{};
  Db.New();
  db := Db.Open(os.Getenv("DB_NOUSER_USER"), os.Getenv("DB_NOUSER_PWD"));
  
  defer db.Close();

  query := "SELECT id FROM `Users` WHERE `Name` LIKE ?";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  row := stmt.QueryRow(u.Name);
  defer stmt.Close();

  var id int;
  row.Scan(&id)

  if(id != 0){return true;}

  return false;
}
