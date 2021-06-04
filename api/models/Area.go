package models

import (
  //"fmt"

	"database/sql"

	db "astara/commons/database"
)

type Area struct {
  //Id int `json:"id"`;
  Name string `json:"name"`;
  Slug string `json:"slug"`;
  id_user int;
  Deleteable int `json:"deleteable"`;
}

type areaDb struct {
  Name sql.NullString;
  Slug sql.NullString;
  Id_user sql.NullInt64;
  Deleteable sql.NullInt64;
}

func GetAreasById (user int, rol string) []Area {
  db := db.GetDb(rol);

  query := "SELECT `Name`,`Slug` FROM `Areas` WHERE `Id_user` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err); }
  defer stmt.Close();

  rows, err := stmt.Query(user);
  if err != nil { panic(err); }

  areas := []Area{};
  area := Area{};
  
  areadb := areaDb{};

  for rows.Next() {

    rows.Scan(&areadb.Name,&areadb.Slug);

    if areadb.Slug.Valid && areadb.Name.Valid {
      area.Name = areadb.Name.String;
      area.Slug = areadb.Slug.String;

      areas = append(areas,area);
    } 
  }

  return areas;
}

func CheckUserArea(user int, slug,rol string) int {
  db := db.GetDb(rol);

  query := "SELECT `Id` FROM `Areas` WHERE `Id_user` LIKE ? AND `Slug` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil { return -1 /*panic(err);*/ }


  row := stmt.QueryRow(user,slug);
  if err != nil { return -1; /*panic(err);*/ }

  defer stmt.Close();
  
  var ( id sql.NullInt64; )

  row.Scan(&id);
  if !id.Valid || id.Int64 == 0 { return  -1; }

  return int(id.Int64);
}

func AreaIsDeleteable(areaId int, rol string) bool {
  db := db.GetDb(rol);

  query := "SELECT `Deleteable` FROM `Areas` WHERE `Id` LIKE ? ;";
  stmt, err := db.Prepare(query);
  if err != nil { return false /*panic(err);*/ }


  row := stmt.QueryRow(areaId);
  if err != nil { return false; /*panic(err);*/ }
  
  defer stmt.Close();

  var ( del sql.NullInt64; )

  row.Scan(&del);
  if !del.Valid { return false }

  switch del.Int64 {
    case 0 : { return false; }
    case 1 : { return true; }
    default: { return false; }
  }
}

func alreadyCreated(uid int, rol, name string) bool {
  db := db.GetDb(rol);

  query := "SELECT `Id` FROM `Areas` WHERE `Name` LIKE ? AND `Id_user` LIKE ?";
  stmt, err := db.Prepare(query);
  if err != nil { /*return false;*/ panic(err); }

  var id  int ;
  row := stmt.QueryRow(name,uid);
  defer stmt.Close();

  row.Scan(&id);
  if id == 0 { return false;
  }else{ return true; }
}

func CreateNewArea(uid int, rol, name, slug string) bool {
  db := db.GetDb(rol);

  if !alreadyCreated(uid,rol,name) {
    query := "INSERT INTO `Areas` (`Name`,`Id_user`,`Deleteable`,`Slug`) VALUES (?, ?, ?, ?)";
    stmt, err := db.Prepare(query);
    if err != nil { /*return false;*/ panic(err); }

    _ , err = stmt.Exec(name,uid,1,slug);
    if err != nil { return false; /*panic(err);*/ }
    defer stmt.Close();
    return true;
  }


  return false;
}

func DelArea(uid int, rol, slug string) bool {
  db := db.GetDb(rol);

  query := "DELETE FROM `Areas` WHERE `Slug` LIKE ? AND `Id_user` LIKE ? AND `Deleteable` = 1;";
  stmt, err := db.Prepare(query);
  if err != nil { return false; /*panic(err);*/ }

  if _, err = stmt.Exec(slug, uid); err != nil {
    return false; /*panic(err);*/
  }else{
    return true;
  }
}

func GetIdFromSlug(uid int, rol, slug string) int {
  db := db.GetDb(rol);

  query := "SELECT `Id`FROM `Areas` WHERE `Slug` LIKE ? AND `Id_user` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil { return -1; /*panic(err);*/ }

  var (
    id sql.NullInt64;
  )
  if err := stmt.QueryRow(slug, uid).Scan(&id); err != nil {
    /*panic(err);*/
    return -1;
  }else{
    if id.Valid {
      return int(id.Int64);
    }
    return -1;
  }
}

func CreateIndexArea(uid int) bool {
  db := db.GetDb("nonuser");

  query := "INSERT INTO `Areas` (`Name`,`Id_user`,`Deleteable`,`Slug`) VALUES (?,?,?,?)";
  stmt, err := db.Prepare(query);
  if err != nil { return false; /*panic(err);*/ }

  if _, err := stmt.Exec("Index", uid, 0, "index"); err != nil { return false /*panic(err);*/ }

  return true;
}
