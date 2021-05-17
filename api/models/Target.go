package models

import (
	"fmt"

	"database/sql"
	"encoding/json"
	//"os"

	db "astara/commons/database"
)

type Target struct {
	Id int `json:"id"`;
	Id_parent int `json:"id_parent"`;
	Id_user int `json:"id_user"`;
	Id_area int `json:"id_area"`;
	Id_status int `json:"id_status"`;
	Name string `json:"name"`;
	Deadline string `json:"deadline"`;
	Children int `json:"children"`;
}

type targetSql struct  {
		id sql.NullInt32;
		id_parent sql.NullInt32;
		id_user sql.NullInt32;
		id_area sql.NullInt32;
		id_status sql.NullInt32;
		name sql.NullString;
		deadline sql.NullString;
		children sql.NullInt32;
	}

func GetTargetsById(user int, rol string) []Target {
	if rol == ""{ return nil; }

	db := db.GetDb(rol);
	query := "SELECT * FROM `Targets` WHERE `Id_usu` LIKE ?";

	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }

	defer stmt.Close();

	rows, err := stmt.Query(user);
	if err != nil { panic(err); }

	targets := []Target{};
	target := Target{};
	t := targetSql{};

	for rows.Next() {

		rows.Scan(&t.id,&t.id_parent,&t.id_user,&t.id_area,&t.id_status,&t.name,&t.deadline,&t.children);

		if t.id.Valid && t.id_user.Valid && t.id_area.Valid && t.id_status.Valid && t.name.Valid && t.deadline.Valid && t.children.Valid {
			if !t.id_parent.Valid {	target.Id_parent = -1; }else{ target.Id_parent = int(t.id_parent.Int32); }

			target.Id = int(t.id.Int32);
			target.Id_user = int(t.id_user.Int32);
			target.Id_area = int(t.id_area.Int32);
			target.Id_status = int(t.id_status.Int32);
			target.Name = t.name.String;
			target.Deadline = t.deadline.String;
			target.Children = int(t.children.Int32);

			targets = append(targets, target);
		}
	}
	return targets;	
}

type UserArea struct {
	Name string `json:"name"`;
	Deadline string `json:"deadline"`;
	Done bool `json:"done"`;
}

type UserAreaDB struct {
	name sql.NullString;
	deadline sql.NullString;
	id_status sql.NullInt32;
}

func GetTargetsUserArea(user, areaid int, rol string) []UserArea {

	db := db.GetDb(rol);
	query := "SELECT `Name`, `Deadline`, `Id_status` FROM `Targets` WHERE `Id_usu` LIKE ? AND `Id_area` LIKE ?";

	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }

	defer stmt.Close();

	rows, err := stmt.Query(user,areaid);
	if err != nil { panic(err); }

	userArea := []UserArea{};
	ua := UserArea{};
	uadb := UserAreaDB{};

	for rows.Next() {

		rows.Scan(&uadb.name,&uadb.deadline,&uadb.id_status);

		if uadb.name.Valid && uadb.deadline.Valid && uadb.id_status.Valid{
			fmt.Println(uadb.id_status.Int32);

			if uadb.id_status.Int32 == 51 { ua.Done = true; }else{ ua.Done = false; }

			ua.Name = uadb.name.String;
			ua.Deadline = uadb.deadline.String;

			userArea = append(userArea,ua);

		}
	}
	return userArea;
}

func targetMarshal(v interface{}) string{
	if json, err := json.Marshal(v); err != nil { 
		//panic(err)
		return "";
	}else{
		return string(json);
	}
}

func GetFormatedUserAreas(user, areaid int, rol string) string {
	return targetMarshal(GetTargetsUserArea(user,areaid, rol));
}
