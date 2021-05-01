package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	m "astara/models"
)

type UserCtr struct {}

func Login(c *fiber.Ctx) error {
	u := m.User{};

	if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}

	if(u.IsRegistered()){
		return c.SendStatus(200);
	}else{
		return c.SendStatus(404);
	}
}

