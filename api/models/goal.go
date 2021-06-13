package models

import (
	"fmt"
	"sort"

	"encoding/json"

	"database/sql"

	db "astara/commons/database"
)

type Goal struct {
	Id            int     `json:"id"`
	Id_target     int     `json:"id_target"`
	Name					string  `json:"name"`
	Deadline			string  `json:"deadline"`
	Description   string  `json:"description"`
	Children			int			`json:"children"`;
	ChildrenDone	int			`json:"childrenDone"`;
	Id_parent			int			`json:"id_parent"`;
	Tasks					*[]Task`json:"tasks"`;
}

func calcProgress(children, childrenDone int64) int {
	return int(childrenDone/children);
}


func GetPaginatedGoals(uid, areaId, size int, rol string, paginated bool) string {
  fmt.Println("Getting paginated tasks of area:");
  db := db.GetDb(rol);

	query := "SELECT TR.`Id`, TR.`Name`, G.`Description`, TR.`Deadline`, G.`Children`, G.`ChildrenDone` FROM `Goals` AS G JOIN `Targets` AS TR ON(G.`Id_target` = TR.`Id`) WHERE TR.`Id_usu` = ? AND TR.`Id_area` = ? AND TR.`Id_status` = 50 ORDER BY TR.`Id` DESC LIMIT ?;";

	stmt, err := db.Prepare(query);
	if err != nil { panic(err); }

	limit := size;
	if paginated {limit += 5;}

	rows, err := stmt.Query(uid, areaId, limit);
	defer stmt.Close();

	var (
		Id, Children, ChildrenDone sql.NullInt64;
		Name, Description, Deadline sql.NullString;
	)

	goals := make(map[int]*Goal);

	for rows.Next(){
		err := rows.Scan(&Id, &Name, &Description, &Deadline, &Children, &ChildrenDone);
		if err != nil { panic(err); }

		//if Id.Valid  || Name.Valid || Description.Valid || Children.Valid || ChildrenDone.Valid {
			goal := Goal{};

			if !Id.Valid { goal.Id = 0; }else{ goal.Id= int(Id.Int64); }
			if !Name.Valid { goal.Name = ""; }else{ goal.Name = Name.String; }
			if !Description.Valid { goal.Description = ""; }else{ goal.Description = Description.String; }
			if !Deadline.Valid { goal.Deadline = ""; }else{ goal.Deadline = Deadline.String; }
			if !Children.Valid { goal.Children = 0; }else{ goal.Children = int(Children.Int64); }
			if !ChildrenDone.Valid { goal.Id = 0; }else{ goal.ChildrenDone = int(ChildrenDone.Int64); }

			goals[int(Id.Int64)] = &goal;
		//}
	}

	//for _, s := range goals{
	//	fmt.Printf("%+v\n",s);
	//}
	

	taskQuery := "SELECT TR.`Id`, TR.`Id_parent`, TR.`Deadline`, TS.`Dated`, TR.`Name` FROM `Task` AS TS JOIN `Targets` AS TR ON (TS.`Id_target` = TR.`Id`) WHERE TR.`Id_usu` = ? AND TR.`Id_area` = ? AND TR.`Id_parent` IS NOT NULL AND TR.`Id_status` = 50;";

	taskStmt, err := db.Prepare(taskQuery);
	if err != nil { panic(err); }

	taskRows, err := taskStmt.Query(uid, areaId);
	if err != nil { panic(err); }
	defer taskStmt.Close();

	var (
		IdTask, Id_parent sql.NullInt64;
		DeadlineTask, Dated, NameTask sql.NullString;
	)

	tasks := []Task{};
	for taskRows.Next() {
		err := taskRows.Scan(&IdTask, &Id_parent, &DeadlineTask, &Dated, &NameTask);
		if err != nil { panic(err); }


		task := Task{};

		if IdTask.Valid { task.Id = int(IdTask.Int64); }else{ task.Id = 0; }
		if Id_parent.Valid { task.Id_parent = int(Id_parent.Int64); }else{ task.Id_parent = 0; }
		if DeadlineTask.Valid { task.Deadline = DeadlineTask.String; }else{ task.Deadline = ""; }
		if NameTask.Valid { task.Name = NameTask.String; }else{ task.Name = ""; }
		if Dated.Valid { task.Dated = Dated.String; }else{ task.Dated = ""; }

		tasks = append(tasks, task);
	}

	fmt.Println("tasks");
	for s := range tasks{
		fmt.Printf("%+v",s);
	}
	
	nestedGoals := nestGoals(tasks, goals);
	//nestedGoals := nestGoals(goals);

	sortedIds:= make([]int, 0 ,len(nestedGoals));
	for k := range goals{
		sortedIds= append(sortedIds, k);
	}

	sort.Ints(sortedIds);

	arrayGoals := arrayGoals(sortedIds, nestedGoals);

	jsonGoals, err := json.Marshal(arrayGoals);
	if err != nil { panic(err); }

	return string(jsonGoals);
}

//delete
/*
func GetAllGoals(uid, areaid int, rol string) string {
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

/*
func nestGoals(goals map[int]*Goal) map[int]*Goal{
	for _, v := range goals {
		if v.Id_parent != 0 {
			if goals[v.Id_parent].Tasks == nil {
				var g = goals[v.Id_parent].Tasks;
				g = &[]Task{*v};
				goals[v.Id_parent].Tasks = g;

				delete(goals,v.Id);
			}
		}
	}
	return goals;
};
*/
func nestGoals(tasks []Task, goals map[int]*Goal) map[int]*Goal {

	for _, v := range tasks{
		if goals[v.Id_parent].Tasks == nil {
			var t = goals[v.Id_parent].Tasks;
			t = &[]Task{v};
			goals[v.Id_parent].Tasks = t;
		}else{
			*goals[v.Id_parent].Tasks = append(*goals[v.Id_parent].Tasks, v);
		}
	}
	return goals;
}

func arrayGoals(order []int, goals map[int]*Goal) []Goal {
	arrayGoals := []Goal{};
	for _,v := range order {
		arrayGoals = append(arrayGoals, *goals[v]);
	}
	return arrayGoals;
}

func CreateNewGoal(uid, areaid int, rol, name, deadline, description string) bool {
  fmt.Println("Create Goal:");
  db := db.GetDb(rol);
  
	targetQuery := "INSERT INTO `Targets` (`Id_usu`,`Id_area`,`Name`,`Deadline`) VALUES(?, ?, ?, ?);";

  targetStmt, err := db.Prepare(targetQuery);
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err); }

	res, err := targetStmt.Exec(uid, areaid, name, deadline);
  defer targetStmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err);}

	targetId, err := res.LastInsertId();
	if err != nil { return false; }

	goalQuery := "INSERT INTO `Goals` (`Id_target`,`Description`) VALUES(?, ?);";

  taskStmt, err := db.Prepare(goalQuery);
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err); }

	_ , err = taskStmt.Exec(targetId,description);
  defer taskStmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/panic(err);}

	return true;
}

func RemoveGoal(uid, id int, rol string) bool {
  fmt.Println("Remove Task:");
  db := db.GetDb(rol);
  
	query := "DELETE FROM `Targets` WHERE `Id` = ? AND `Id_usu` = ?;";

  stmt, err := db.Prepare(query);
  if err != nil && err != sql.ErrNoRows { return false; /*panic(err);*/ }

	_, err = stmt.Exec(id, uid);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { return false; /*panic(err);*/}

	return true;
}

func UpdateExistingGoal(uid, taskId int, rol, name, description string) bool {
  fmt.Println("Update Task:");
  db := db.GetDb(rol);
  
	queryTarget := "UPDATE `Targets` SET `Name ` = ? WHERE Id = ?;";

  stmt, err := db.Prepare(queryTarget);
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err); }

	_, err = stmt.Exec(name, taskId);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err);}

	queryTask := "UPDATE `Goal` SET Description = ? WHERE Id_target = ?;";

  stmt, err = db.Prepare(queryTask);
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err); }

	_, err = stmt.Exec(description, taskId);
  defer stmt.Close();
  if err != nil && err != sql.ErrNoRows { /*return false;*/ panic(err);}

	return true;
}
