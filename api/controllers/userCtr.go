package controllers

import (
	"fmt"
	"reflect"

	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	"astara/commons/cookie"
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
	fmt.Println("Check");
	//other way to do it
	type response struct{
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil{return err;}

	id := CheckCredentials(res.User, res.Pass);

	c.Status(200);

	if id == -1 { return c.JSON(fiber.Map{"logged":"false"}); }

	token := jwt.CreateToken(id); // Create a token
	cookie := cookie.CreateCookie(token); // Create a cookie with that token value
	fmt.Println(reflect.TypeOf(cookie));
	fmt.Println(cookie.Value);
	c.Cookie(cookie); // Set the cookie

	return c.JSON(fiber.Map{"logged":"true"});
}
