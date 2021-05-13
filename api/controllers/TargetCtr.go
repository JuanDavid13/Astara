package controllers 

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

func GetTargets(c *fiber.Ctx) error {

	user := jwt.GetUser(c.Cookies("token"));
	targets := GetTargetsById(user)

	Targets, err := json.Marshal(targets);
	if err != nil { panic(err); }

	c.Status(fiber.StatusOK);
	return c.JSON(string(Targets));
}
