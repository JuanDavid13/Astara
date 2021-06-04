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
		Slug string `json:"slug"`;
		Name string `json:"name"`;
		Deadline string `json:"deadline"`;
		Dated string `json:"dated"`;
	}
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil { panic(err); }

	cl := c.Locals("claims").(jwt.Claims);

	id := GetIdFromSlug(cl.User, cl.Rol, res.Slug);
	if id == -1 {
		return c.JSON(fiber.Map{
			"created":false,
		});
	}

	created := CreateNewTask(cl.User, id, cl.Rol, res.Name, res.Deadline, res.Dated);
	if !created {
		return c.JSON(fiber.Map{
			"created":false,
		});
	}
	return c.JSON(fiber.Map{
		"created":true,
	});
}
