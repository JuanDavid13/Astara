package controllers

import (
	"fmt"
	"encoding/json"

	"astara/commons/jwt"
	. "astara/models"

	"github.com/gofiber/fiber/v2"
)


func GetGoals(c *fiber.Ctx) error {
	fmt.Println("Get goals:")

	cl := c.Locals("claims").(jwt.Claims)

	goals := GetGoalsbyId(cl.User,cl.Rol);

	fmt.Printf("%+v\n",goals);

	Goals, err := json.Marshal(goals);
	if err != nil { panic(err); }

	c.Status(200)
	return c.JSON(string(Goals));

}
