package models

import (
	"fmt"
	"sort"
	//"reflect"
	//"os"
	"encoding/json"

	db "astara/commons/database"
	"database/sql"
)

type Task struct{
	Id int `json:"id"`;
	Id_parent int `json:"id_parent"`;
	Id_usu int `json:"id_usus"`;
	Id_area int `json:"id_area"`;
	Id_status int `json:"id_status"`;
	Children int `json:"children"`;
	ChildrenDone int `json:"chilldrenDone"`;
	Name string `json:"name"`;
	Deadline string `json:"deadline"`;
	Dated string `json:"dated"`;
	Tasks *[]Task `json:"tasks"`;
}

func GetPaginatedTasksOfArea(uid, areaId, offset int, rol string) string {
  fmt.Println("Getting paginated tasks of area:");
  db := db.GetDb(rol);

	query := "SELECT TR.`Id`, TR.`Id_parent`, TR.`Name`, TR.`Deadline`,TR.`Children`,TR.`ChildrenDone`, TS.`Dated` FROM `Targets` AS TR JOIN `Task` AS TS ON(TR.`Id` = TS.`Id_target`) WHERE TR.`Id_usu` = ? AND TR.`Id_area` = ? AND TR.`Id_status` = ? ORDER BY TR.`Id` DESC LIMIT ?, 7;";

	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }

	rows, err := stmt.Query(uid,areaId,50, offset);
	defer stmt.Close();

	var (
		Id, Id_parent, Children, ChildrenDone sql.NullInt64;
		Name, Deadline, Dated sql.NullString;
	)

	tasks := make(map[int]*Task);

	for rows.Next(){
		err := rows.Scan(&Id, &Id_parent, &Name, &Deadline, &Children, &ChildrenDone, &Dated);
		if err != nil { panic(err); }

		if Id.Valid  || Name.Valid || Deadline.Valid || Children.Valid || ChildrenDone.Valid {
			task := Task{};

			if !Id.Valid { task.Id = 0; }else{ fmt.Println(Id.Int64);task.Id= int(Id.Int64); }
			if !Name.Valid { task.Name = ""; }else{ task.Name = Name.String; }
			if !Deadline.Valid { task.Deadline = ""; }else{ task.Deadline = Deadline.String; }
			if !Dated.Valid { task.Deadline = ""; }else{ task.Dated = Dated.String; }
			if !Children.Valid { task.Children = 0; }else{ task.Children = int(Children.Int64); }
			if !ChildrenDone.Valid { task.Id = 0; }else{ task.ChildrenDone = int(ChildrenDone.Int64); }
			if !Id_parent.Valid { task.Id_parent = 0; }else{ task.Id_parent = int(Id_parent.Int64); }

			tasks[int(Id.Int64)] = &task;
		}
	}
	for _, s := range tasks{
		fmt.Printf("%+v\n",s);
	}
	
	nestedTasks := nestTasks(tasks);

	sortedIds:= make([]int, 0 ,len(nestedTasks));
	for k := range nestedTasks {
		sortedIds= append(sortedIds, k);
	}

	sort.Ints(sortedIds);

	arrayTasks := arrayTasks(sortedIds, nestedTasks);

	jsonTasks, err := json.Marshal(arrayTasks);
	if err != nil { panic(err); }

	return string(jsonTasks);

}
//delete
/*
func GetAllTasksOfArea(uid, areaid int, rol string) string {
  fmt.Println("Getting all tasks of area:");
  db := db.GetDb(rol);

	query := "SELECT TR.`Id`, TR.`Id_parent`, TR.`Name`, TR.`Deadline`,TR.`Children`,TR.`ChildrenDone`, TS.`Dated` FROM `Targets` AS TR JOIN `Task` AS TS ON(TR.`Id` = TS.`Id_target`) WHERE TR.`Id_usu` = ? AND TR.`Id_area` = ? AND TR.`Id_status` = ? ORDER BY TR.`Id` DESC;";

	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }

	rows, err := stmt.Query(uid,areaid,50);
	defer stmt.Close();

	var (
		Id, Id_parent, Children, ChildrenDone sql.NullInt64;
		Name, Deadline, Dated sql.NullString;
	)

	tasks := make(map[int]*Task);

	for rows.Next(){
		err := rows.Scan(&Id, &Id_parent, &Name, &Deadline, &Children, &ChildrenDone, &Dated);
		if err != nil { panic(err); }

		if Id.Valid  || Name.Valid || Deadline.Valid || Children.Valid || ChildrenDone.Valid {
			task := Task{};

			if !Id.Valid { task.Id = 0; }else{ fmt.Println(Id.Int64);task.Id= int(Id.Int64); }
			if !Name.Valid { task.Name = ""; }else{ task.Name = Name.String; }
			if !Deadline.Valid { task.Deadline = ""; }else{ task.Deadline = Deadline.String; }
			if !Dated.Valid { task.Deadline = ""; }else{ task.Dated = Dated.String; }
			if !Children.Valid { task.Children = 0; }else{ task.Children = int(Children.Int64); }
			if !ChildrenDone.Valid { task.Id = 0; }else{ task.ChildrenDone = int(ChildrenDone.Int64); }
			if !Id_parent.Valid { task.Id_parent = 0; }else{ task.Id_parent = int(Id_parent.Int64); }

			tasks[int(Id.Int64)] = &task;
		}
	}
	for _, s := range tasks{
		fmt.Printf("%+v\n",s);
	}
	
	//nestedTasks := nestTasks(tasks);

	//arrayTasks := arrayTasks(nestedTasks);

	jsonTasks, err := json.Marshal(arrayTasks);
	if err != nil { panic(err); }

	return string(jsonTasks);
}
*/

func nestTasks(tasks map[int]*Task) map[int]*Task{
	for _, v := range tasks	{
		if v.Id_parent != 0 {
			if tasks[v.Id_parent].Tasks == nil {
				var t = tasks[v.Id_parent].Tasks;
				t = &[]Task{*v};
				tasks[v.Id_parent].Tasks = t;

				delete(tasks,v.Id);
			}
		}
	}
	return tasks;
};

//func arrayTasks(tasks map[int]*Task) []Task{
func arrayTasks(order []int, tasks map[int]*Task) []Task{
	arrayTasks := []Task{};
	//for _,v :=range tasks {
	for _,v :=range order {
		//arrayTasks = append(arrayTasks,*v);
		arrayTasks = append(arrayTasks,*tasks[v]);
	}
	return arrayTasks;
}

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

func RemoveTask(uid, id int, rol string) bool {
  fmt.Println("Remove Task:");
  db := db.GetDb(rol);
  
	query := "DELETE FROM `Targets` WHERE `Id` = ? AND `Id_usu` = ?;"; //double check for security

  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return false; /*panic(err);*/ }

	_, err = stmt.Exec(id, uid);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return false; /*panic(err);*/}

	return true;
}

func UpdateTask(uid, taskId int, rol, name, deadline, dated string) bool {
  fmt.Println("Update Task:");
  db := db.GetDb(rol);
  
	queryTarget := "UPDATE `Targets` SET Name = ?, Deadline = ? WHERE Id = ?;";

  stmt, err := db.Prepare(queryTarget);
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err); }

	_, err = stmt.Exec(name, deadline, taskId);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err);}

	queryTask := "UPDATE `Task` SET Dated = ? WHERE Id_target = ?;";

  stmt, err = db.Prepare(queryTask);
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err); }

	_, err = stmt.Exec(dated, taskId);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err);}

	return true;
}
