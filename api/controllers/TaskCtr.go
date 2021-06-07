package controllers

import (
	//"fmt"
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
	}
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { /*return c.SendStatus(400);*/ panic(err); }

	if res.Name == "" || res.Deadline == "" || res.Slug == "" { return c.SendStatus(400); }


	cl := c.Locals("claims").(jwt.Claims);

	//get the id of the area it is from 
	id := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if id == -1 { return c.JSON(fiber.Map{ "created":false, }); }

	//create a new task
	if !CreateNewTask(cl.User, id, cl.Rol, res.Name, res.Deadline, res.Dated) {
		return c.JSON(fiber.Map{ "created":false, });
	}else{
		return c.JSON(fiber.Map{ "created":true, });
	}
}

func GetAllTasks(c *fiber.Ctx) error {
	type response struct { Slug string `json:"slug"`; }
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { return c.SendStatus(400); /*panic(err);*/ }

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
