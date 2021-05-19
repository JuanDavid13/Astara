package models

import (
	"fmt"

	"database/sql"
	//"encoding/json"

	db "astara/commons/database"
)

type Goal struct {
	Id             int     `json:"id"`
	Id_target      int     `json:"id_target"`
	Name					 string  `json:"name"`
	Deadline			 string  `json:"deadline"`
	Description    string  `json:"description"`
	Quantity       float64 `json:"quantity"`
	ActualQuantity float64 `json:"actualQuantity"`
	Progress			 int		 `json:"progress"`
}

func calcProgress(children, childrenDone int64, quant, actualQuant float64) int {
	progress := 0;
	if children == 0 {
		progress = int(actualQuant/quant);
	}else{
		qua := int(actualQuant/quant);
		chi := int(childrenDone/children);
		progress = (qua + chi)/2
	}
	return progress
}


func GetGoalsbyId(user int, rol string) []Goal {
	if rol == "" { return nil; }

	type goalSQL struct {
		Id sql.NullInt64;
		Name sql.NullString
		Deadline sql.NullString
		Children sql.NullInt64;
		ChildrenDone sql.NullInt64;
		Quantity sql.NullFloat64
		ActualQuantity sql.NullFloat64
	}

  goals := []Goal{};
  goal := Goal{};
  goalsql := goalSQL{};

	db := db.GetDb(rol);
	query := "SELECT T.`Id`, T.`Name`, T.`Deadline`, T.`Children`,T.`ChildrenDone`, G.`Quantity`,G.`ActualQuant` FROM `Goals` AS G JOIN `Targets` AS T ON (G.`Id_target` = T.`Id`) WHERE T.`Id_usu` LIKE ? AND T.`Id_parent`IS NULL AND T.`Id_status` = 50";


	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }
	defer stmt.Close();

	rows, err := stmt.Query(user);
	if err != nil { panic(err); }

	for rows.Next() {

		err = rows.Scan(&goalsql.Id, &goalsql.Name, &goalsql.Deadline, &goalsql.Children, &goalsql.ChildrenDone, &goalsql.Quantity, &goalsql.ActualQuantity);
		if err != nil { panic(err); }

			fmt.Printf("%+v\n",rows);

		if goalsql.Id.Valid && goalsql.Name.Valid && goalsql.Deadline.Valid && goalsql.Children.Valid && goalsql.ChildrenDone.Valid && goalsql.Quantity.Valid && goalsql.ActualQuantity.Valid {

			fmt.Println("entra");

			goal.Id = int(goalsql.Id.Int64);
			goal.Name = goalsql.Name.String;
			goal.Deadline = goalsql.Deadline.String;
			progress := calcProgress(goalsql.Children.Int64,goalsql.ChildrenDone.Int64, goalsql.Quantity.Float64, goalsql.ActualQuantity.Float64)
			goal.Progress = progress;

			goals = append(goals, goal);
		}
	}
	return goals;
}
