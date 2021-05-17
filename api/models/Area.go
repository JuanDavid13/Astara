package models

import (
  "fmt"

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
  fmt.Println("rol");
  fmt.Println(rol);
  db := db.GetDb(rol);

  fmt.Println("en area");
  fmt.Println(db);

  query := "SELECT `Name`,`Slug`,`Deleteable` FROM `Areas` WHERE `Id_user` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil { panic(err); }

  defer stmt.Close();

  rows, err := stmt.Query(user);
  if err != nil { panic(err); }

  areas := []Area{};
  area := Area{};
  
  
  areadb := areaDb{};

  for rows.Next() {

    rows.Scan(&areadb.Name,&areadb.Slug,&areadb.Deleteable);

    if areadb.Slug.Valid && areadb.Name.Valid && areadb.Deleteable.Valid {
      area.Name = areadb.Name.String;
      area.Slug = areadb.Slug.String;
      area.Deleteable = int(areadb.Deleteable.Int64);

      areas = append(areas,area);
    } 
  }
  return areas;
}

func CheckUserArea(user int, slug,rol string) (string,bool){
  db := db.GetDb(rol);

  query := "SELECT `Id` FROM `Areas` WHERE `Id_user` LIKE ? AND `Slug` LIKE ?;";
  stmt, err := db.Prepare(query);
  if err != nil { panic(err); }

  defer stmt.Close();

  row := stmt.QueryRow(user,slug);
  if err != nil { panic(err); }
  
  var id int;
  row.Scan(&id);
  if id != 0 { return GetFormatedUserAreas(user,id,rol),true;  }

  return "",false;
}

