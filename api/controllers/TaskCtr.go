package controllers

import (
	//"fmt"
	"strconv"
	//"reflect"

	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)


func CreateTask(c *fiber.Ctx) error {
	type response struct {
		Slug string `json:"slug"`; // to distinc the area it is from
		Name string `json:"name"`;
		Deadline string `json:"deadline"`;
		Dated string `json:"dated"`;
		Id int `json:"id"`;
	}
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { /*return c.SendStatus(400);*/ panic(err); }

	if res.Name == "" || res.Deadline == "" || res.Slug == "" || res.Id == 0 { return c.SendStatus(400); }


	cl := c.Locals("claims").(jwt.Claims);

	//get the id of the area it is from 
	areaId := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if areaId == -1 { return c.JSON(fiber.Map{ "created":false, }); }

	//create a new task
	if !CreateNewTask(cl.User, areaId, res.Id, cl.Rol, res.Name, res.Deadline, res.Dated) {
		return c.JSON(fiber.Map{ "created":false, });
	}else{
		return c.JSON(fiber.Map{ "created":true, });
	}
}

func GetPagTasks(c *fiber.Ctx) error {
	slug := c.Params("slug");
	size := c.Params("size");
	pag:= c.Params("paginated");

	paginated, err := strconv.ParseBool(pag);
	if err != nil { return c.SendStatus(400); }

	sizeInt, err := strconv.Atoi(size);
	if err != nil { /*panic(err);*/ return c.SendStatus(400); }

	if sizeInt < 0  || slug == "" { return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);
	
	id := GetIdFromSlug(cl.User, cl.Rol, slug);
	if id == -1 { return c.JSON(fiber.Map{ "error":true, }); }

	tasks := GetPaginatedTasks(cl.User, id, sizeInt, cl.Rol, paginated);

	return c.JSON(fiber.Map{
		"error":false,
		"tasks":tasks,
	});

}

//delete
/*
func GetPaginatedTasks(c *fiber.Ctx) error {
	type response struct {
		Slug string `json:"slug"`;
		Offset int `json:"offset"`;
	}
	res := response{};

	fmt.Println(res.Offset);
	fmt.Println(res.Slug);

	if err := json.Unmarshal(c.Body(), &res); err != nil { panic(err); /*return c.SendStatus(400); }
	if res.Offset < 0  || res.Slug == "" { return c.SendStatus(404); }

	cl := c.Locals("claims").(jwt.Claims);
	
	id := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if id == -1 { return c.JSON(fiber.Map{ "error":true, }); }

	tasks := GetPaginatedTasksOfArea(cl.User, id, res.Offset, cl.Rol);

	return c.JSON(fiber.Map{
		"error":false,
		"tasks":tasks,
	});

}
*/

//delete
/*
func GetAllTasks(c *fiber.Ctx) error {
	type response struct { Slug string `json:"slug"`; }
	res := response{};

	//if err := json.Unmarshal(c.Body(),&res); err != nil { return c.SendStatus(400); /*panic(err); }

	if res.Slug == "" { return c.SendStatus(400); }


	cl := c.Locals("claims").(jwt.Claims);

	//get the id of the area it is from 
	id := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if id == -1 { return c.JSON(fiber.Map{ "error":true, }); }

	//create a new task
	tasks := GetAllTasksOfArea(cl.User,id,cl.Rol);

	return c.JSON(fiber.Map{
		"tasks":tasks,
	});
}
*/

func DeleteTask(c *fiber.Ctx) error {
	
	type response struct { Id int `json:"id"`; }
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { return c.SendStatus(400); /*panic(err);*/ }

	if res.Id == 0 { return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);

	c.Status(200);
	if !RemoveTask(cl.User, res.Id, cl.Rol){
		return c.JSON(fiber.Map{ "deleted":false, });
	}
	return c.JSON(fiber.Map{ "deleted":true, });
}

func EditTask(c *fiber.Ctx) error {
	type response struct {
		Name string `json:"name"`;
		Deadline string `json:"deadline"`;
		Dated string `json:"dated"`;
		TaskId int `json:"task_id"`;
	}	
	res := response{};

	if err := json.Unmarshal(c.Body(), &res); err != nil { panic(err); }

	if res.Name == "" || res.Deadline == "" || res.Dated == "" || res.TaskId <= 0{ return c.SendStatus(400); }
	
	cl := c.Locals("claims").(jwt.Claims);


	if !UpdateTask(cl.User, res.TaskId, cl.Rol, res.Name, res.Deadline, res.Dated) {
		return c.JSON(fiber.Map{ "updated":false });
	}else{
		return c.JSON(fiber.Map{ "updated":true });
	}
}
