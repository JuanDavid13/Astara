package controllers

import (
	"fmt"
	//"reflect"

	"encoding/json"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	"astara/commons/cookie"
	. "astara/models"
)

//sanitize
func CheckUser(c *fiber.Ctx) error {
	fmt.Println("Check user");
	u := User{};
	if err := json.Unmarshal(c.Body(),&u); err != nil{ return c.SendStatus(400); }

	c.Status(200);
	if IsRegistered(u.User){ return c.JSON(fiber.Map{"found":"true",})
	}else{ return c.JSON(fiber.Map{"found":"false",})}
}

func logUser(c *fiber.Ctx, name string, id int, rol string) {
	token := jwt.CreateToken(id,rol); // Create a token
	cookie := cookie.CreateCookie(token); // Create a cookie with that token value
	c.Cookie(cookie); // Set the cookie
}

func Check(c *fiber.Ctx) error {
	fmt.Println("Check");
	//other way to do it
	type response struct{
		User string `json:"user"`
		Pass string `json:"pass"`
	};
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil{/*return err;*/ c.SendStatus(400);}

	c.Status(200);
	if id, correct := CheckCredentials(res.User, res.Pass); !correct {
		return c.JSON(fiber.Map{"logged":"false"});
	}else{
		name, rol := GetBasicUserInfo(id);
		if name == "" || rol == "" { return c.SendStatus(400); }
		logUser(c, name, id, rol);
		return c.JSON(fiber.Map{"logged":"true"});
	}
}

func LogOut(c *fiber.Ctx) error {
	if !cookie.CheckIsEmpty(c.Cookies("token")) {
		cookie := cookie.CreateExpiredCookie();
		c.Cookie(cookie);
		return c.SendStatus(200);
	}
	return c.SendStatus(400);
}

func CreateUser(c *fiber.Ctx) error {
	type response struct {
		User string `json:"user"`;
		Pass string `json:"pass"`;
		Email string `json:"email"`;
	}
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil {
		panic(err);
	}

	if IsRegisteredWithEmail(res.User, res.Email) {
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": true,
			"alreadyCreated": true,
			"created": false,
		});
	}
	
	id, created := CreateNewUser(res.User,res.Pass, res.Email);
	if !created {
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": true,
			"alreadyCreated": false,
			"created": false,
		});
	}else{
		logUser(c, res.User, id, "user");
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": false,
			"alreadyCreated": false,
			"created": true,
		});

	}
}
