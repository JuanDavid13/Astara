package controllers 

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

func GetTargets(c *fiber.Ctx) error {

	cl := c.Locals("claims").(jwt.Claims);
	targets := GetTargetsById(cl.User);

	Targets, err := json.Marshal(targets);
	if err != nil { panic(err); }

	c.Status(200);
	return c.JSON(string(Targets));
}
