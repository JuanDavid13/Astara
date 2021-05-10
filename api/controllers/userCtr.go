package controllers

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

//sanitize
func CheckUser(c *fiber.Ctx) error {
	u := User{};
	if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}

	c.Status(200);

	if IsRegistered(u.User){
		return c.JSON(fiber.Map{"found":"true",})
	}else{
		return c.JSON(fiber.Map{"found":"false",})}
}

func Check(c *fiber.Ctx) error {
	type response struct{
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil{return err;}

	id := CheckCredentials(res.User, res.Pass);

	if( id != -1){
		token := jwt.CreateToken(id);

		cookie := new(fiber.Cookie);
		cookie.Name = "token";
		cookie.Value = token;
		cookie.Expires = time.Now().Add(time.Minute*10);
		cookie.HTTPOnly = true;
		//cookie.Secure = false; //por ahora

		c.Cookie(cookie);

		c.Status(200);
		return c.JSON(fiber.Map{"logged":"true"});
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{"logged":"false"});
	}
}
