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

	//get the id of the area it is from 
	id := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if id == -1 { return c.JSON(fiber.Map{ "created":false, }); }

	//create a new task
	if !CreateNewGoal(cl.User, id, cl.Rol, res.Name, res.Deadline, res.Description) {
		return c.JSON(fiber.Map{ "created":false, });
	}else{
		return c.JSON(fiber.Map{ "created":true, });
	}
}

func GetPaginatedGoals(c *fiber.Ctx) error {
	//type response struct {
	//	Slug string `json:"slug"`; //--> llega como parametro
	//	Offset int `json:"offset"`; //--> llega como parametro
	//	//cambiar offset por limit --> primera solucion
	//}
	//
	//res := response{};

	//fmt.Println(res.Offset);
	//fmt.Println(res.Slug);

	//if err := json.Unmarshal(c.Body(), &res); err != nil { panic(err); /*return c.SendStatus(400);*/ }

	slug := c.Params("slug");
	offset := c.Params("offset");

	offsetInt, err := strconv.Atoi(offset);
	if err != nil { panic(err); /*return c.SendStatus(400);*/ }

	if offsetInt < 0  || slug == "" { return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);
	
	id := GetIdFromSlug(cl.User, cl.Rol, slug);
	if id == -1 { return c.JSON(fiber.Map{ "error":true, }); }

	tasks := GetPaginated(cl.User, id, offsetInt, cl.Rol);

	return c.JSON(fiber.Map{
		"error":false,
		"tasks":tasks,
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
		GoalId int `json:"goalId"`;
	}	
	res := response{};

	if err := json.Unmarshal(c.Body(), &res); err != nil { /*return c.SendStatus(400);*/ panic(err); }

	if res.Name == "" || res.Description == "" || res.GoalId <= 0{ return c.SendStatus(400); }
	
	cl := c.Locals("claims").(jwt.Claims);

	if !UpdateExistingGoal(cl.User, res.GoalId, cl.Rol, res.Name, res.Description) {
		return c.JSON(fiber.Map{ "updated":false });
	}else{
		return c.JSON(fiber.Map{ "updated":true }); }
}
