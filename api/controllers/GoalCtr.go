package controllers

import (
	//"fmt"
	"strconv"
	//reflect

	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

func CreateGoal(c *fiber.Ctx) error {

	type response struct {
		Slug string `json:"slug"`; // to distinc the area it is from
		Name string `json:"name"`;
		Deadline string `json:"deadline"`;
		Description string `json:"description"`;
	}
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { /*return c.SendStatus(400);*/ panic(err); }

	if res.Name == "" || res.Deadline == "" || res.Slug == "" { return c.SendStatus(400); }


	cl := c.Locals("claims").(jwt.Claims);

	areaid := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if areaid == -1 { return c.JSON(fiber.Map{ "error":true, }); }

	if !CreateNewGoal(cl.User, areaid, cl.Rol, res.Name, res.Deadline, res.Description) {
		return c.JSON(fiber.Map{ "error":true, });
	}else{
		return c.JSON(fiber.Map{ "error":false, });
	}
}

func GetPagGoals(c *fiber.Ctx) error {
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

	goals := GetPaginatedGoals(cl.User, id, sizeInt, cl.Rol, paginated);

	return c.JSON(fiber.Map{
		"error":false,
		"goals":goals,
	});

}

//delete
/*
func GetAllGoals(c *fiber.Ctx) error {
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

func DeleteGoal(c *fiber.Ctx) error {
	
	type response struct { Id int `json:"id"`; }
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { return c.SendStatus(400); /*panic(err);*/ }

	if res.Id == 0 { return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);

	c.Status(200);
	if !RemoveGoal(cl.User, res.Id, cl.Rol){
		return c.JSON(fiber.Map{ "deleted":false, });
	}

	return c.JSON(fiber.Map{ "deleted":true, });
}

func EditGoal(c *fiber.Ctx) error {

	type response struct {
		Name string `json:"name"`;
		Description string `json:"description"`;
		Deadline string `json:"deadline"`;
		GoalId int `json:"goal_id"`;
	}	
	res := response{};

	if err := json.Unmarshal(c.Body(), &res); err != nil { /*return c.SendStatus(400);*/ panic(err); }

	if res.Name == "" || res.GoalId <= 0{ return c.SendStatus(400); }
	
	cl := c.Locals("claims").(jwt.Claims);

	if !UpdateGoal(cl.User, res.GoalId, cl.Rol, res.Name, res.Description,res.Deadline) {
		return c.JSON(fiber.Map{ "updated":false });
	}else{
		return c.JSON(fiber.Map{ "updated":true }); }
}
