package models

import (
	"database/sql"
	"os"

	db "astara/commons/database"
	//jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
  Id int `json:"id"`;
  User string `json:"user"`;
  Pass string `json:"pwd"`;
  //Token jwt.Token `json:token`;
}

//func Authorice (name, pwd string) *jwt.Token {
    
//}

func IsRegistered(user string) bool {
  Db := db.Db{};
  Db.New();
  db := Db.Open(os.Getenv("DB_NOUSER_USER"), os.Getenv("DB_NOUSER_PWD"));
  
  defer db.Close();

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
  Db := db.Db{};
  Db.New();
  db := Db.Open(os.Getenv("DB_NOUSER_USER"), os.Getenv("DB_NOUSER_PWD"));
  
  defer db.Close();

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




  type Area struct{
    Name string `json:"name"`;
    Deleteable int `json:"deleteable"`;
  }

func GetUserAreas(user int) []Area {
  Db := db.Db{};
  Db.New();
  db := Db.Open(os.Getenv("DB_USER_USER"), os.Getenv("DB_USER_PWD"));
  
  defer db.Close();

  query := "SELECT  `Name`,`Deleteable` FROM `Areas` WHERE (`Id_user` LIKE ?);";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  rows,err := stmt.Query(user);
  if err != nil{panic(err);}
  defer stmt.Close();

  area := Area{};

  var (
    name sql.NullString;
    deleteable sql.NullInt64;
    Areas []Area;
  )

  for rows.Next() {
    rows.Scan(&name,&deleteable)
    if name.Valid { area.Name = name.String; }
    if deleteable.Valid { area.Deleteable = int(deleteable.Int64); }

    Areas = append(Areas,area);
  }
  return Areas;
}
