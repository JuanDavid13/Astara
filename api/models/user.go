package models

import (
	"fmt"
	//"reflect"
	//"os"

	db "astara/commons/database"
	"database/sql"
  "golang.org/x/crypto/bcrypt"
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
  if err != nil{ panic(err); }

  var id int;
  err = stmt.QueryRow(user, user).Scan(&id);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return false; }
  if id != 0 {return true;}
  return false;
}
func IsRegisteredWithEmail(user, email string) bool {
  db := db.GetDb("nonuser");
  query := "SELECT `Id` FROM Users WHERE  Name LIKE ? OR Email LIKE ?;";

  stmt, err := db.Prepare(query);
  if err != nil{ panic(err); }

  var id int;
  err = stmt.QueryRow(user, email).Scan(&id);
  defer stmt.Close();
  fmt.Println(id);
  if err != nil && err != sql.ErrNoRows { return false; }
  if id != 0 {return true;}

  return false;
}

func CheckCredentials(user, pass string) (int,bool) {
  fmt.Println("is registered:")

  db := db.GetDb("nonuser");
  //query := "SELECT U.`Id`, R.`Name` FROM `Users` AS U JOIN `Rols` AS R ON (U.`Id_rol` = R.`Id` ) WHERE (U.`Name` LIKE ? OR U.`Email` LIKE ?) AND U.`Password` LIKE ?";
  query := "SELECT `Id`, `Password` FROM `Users` WHERE `Name` LIKE ? OR `Email` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil{panic(err);}

  var ( 
    id int;
    pwd string; 
  )
  //err = stmt.QueryRow(user,user,pass).Scan(&id,&rol);
  err = stmt.QueryRow(user,user).Scan(&id,&pwd);
  defer stmt.Close();

  if err != nil && err != sql.ErrNoRows { return -1,false; }

  if err := bcrypt.CompareHashAndPassword([]byte(pwd),[]byte(pass)); err != nil { return -1,false;/*panic(err);*/ }

  return id,true;
}

func createPass(pass string) string {
  encpass, err := bcrypt.GenerateFromPassword([]byte(pass),10);
  if err != nil { panic(err); }
  return string(encpass);
}

func CreateNewUser(user, pass, email string) (int, bool) {
  db := db.GetDb("nonuser");
  query := "INSERT INTO Users (`Name`,`Password`,`Email`) VALUES(?,?,?);";

  encPass := createPass(pass);

  stmt, err := db.Prepare(query);
  if err != nil { panic(err); /*return -1,false;*/ }

  res, err := stmt.Exec(user,encPass,email);
  if err != nil { panic(err); /*return -1,false;*/ }
  id, err := res.LastInsertId();
  if err != nil { panic(err); }

  return int(id),true;
}

func GetBasicUserInfo(user int) (string,string) {
  fmt.Println("get basic user info:")
  db := db.GetDb("nonuser");
  query := "SELECT U.`Name`, R.`Name` FROM `Users` AS U JOIN `Rols`AS R ON (U.`Id_rol` = R.`Id`) WHERE U.`Id` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return "","";/*panic(err);*/ }

  var ( name, rol string; )

  err = stmt.QueryRow(user).Scan(&name,&rol);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return "","";/*panic(err);*/ }

  fmt.Println("name");
  fmt.Println(name);
  fmt.Println(rol);

  return name,rol;
}

func GetBasicInfo(user int, rol string) (string,string,bool,bool){
  fmt.Println("get info:")
  db := db.GetDb(rol);
  query := "SELECT `Name`, `Email`, `Theme` FROM `Users` WHERE `Id` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return "","",false,true;/*panic(err);*/ }

  var (
    name, email sql.NullString;
    theme sql.NullBool;
  )

  err = stmt.QueryRow(user).Scan(&name,&email,&theme);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return "","",false,true;/*panic(err);*/ }

  if name.Valid && email.Valid && theme.Valid {
    return name.String,email.String, theme.Bool, false
  }

  return "","",false,true;
}

func ComparePass(user int, rol, pass string) (bool,bool){
  fmt.Println("compare pass:");
  fmt.Println(pass);
  db := db.GetDb(rol);
  query := "SELECT `password` FROM `Users` WHERE `Id` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return false,true /*panic(err);*/ }

  var pwd sql.NullString;

  err = stmt.QueryRow(user).Scan(&pwd);
  defer stmt.Close();

  if err != nil && err != sql.ErrNoRows { return false,true;/*panic(err);*/ }

  if pwd.Valid {
    if err = bcrypt.CompareHashAndPassword([]byte(pwd.String),[]byte(pass)); err != nil {
      return false,false;
    }else{
      return true,false;
    }
  }

  return false,true;
}

func UpdateUserInfo(uid int, rol, username, email string, theme bool ) bool {
  fmt.Println("Update user info:");
  db := db.GetDb(rol);
  
  query := "UPDATE `Users` SET `Name` = ?, `Email` = ? WHERE Id LIKE ?;";

  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return false;/*panic(err);*/ }

  _, err = stmt.Exec(username, email, uid);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return false;/*panic(err);*/
  }else{ return true; }
}
