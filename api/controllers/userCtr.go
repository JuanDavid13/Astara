package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

type UserCtr struct {}

func IsRegistered(c *fiber.Ctx) error {
	u := User{};

	if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}

	if(u.IsRegistered()){
		c.Status(200);
		return c.JSON(fiber.Map{
			"found":"true",
		});
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{
			"found":"false",
		});
	}
}

func CheckCredentials(c *fiber.Ctx) error {
	u := User{};

	if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}
	id := u.CheckCredentials()
	if( id != -1){
		token := jwt.CreateToken(id);

		cookie := new(fiber.Cookie);
		cookie.Name = "token";
		cookie.Value = token;
		cookie.Expires = time.Now().Add(time.Minute*10);
		cookie.HTTPOnly = true;
		cookie.Secure = false; //por ahora

		c.Cookie(cookie);

		c.Status(200);
		return c.JSON(fiber.Map{
			"logged":"true",
		});
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{
			"logged":"false",
		});
	}
	
	//type Response struct{
	//	Token string `json:"token"`;
	//}
	//response := Response{};
	//c.BodyParser(&response);
}

func CheckToken(c *fiber.Ctx) error {
	//fmt.Println(c.Get("token")); //token
	//fmt.Println(c.Get("Authorization")); // Bearer token
	fmt.Println(c.Cookies("token")); //token

	return c.SendStatus(200);
}

