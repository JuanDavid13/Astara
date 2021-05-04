package controllers

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

func CheckUser(c *fiber.Ctx) error {
	//sanitize
	u := User{};
	if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}

	c.Status(200);

	if IsRegistered(u.User){
		return c.JSON(fiber.Map{
			"found":"true",
		})
	}else{
		return c.JSON(fiber.Map{
			"found":"false",
		})
	}

}

func Check(c *fiber.Ctx) error {

	//u := User{};
	//tambien se puede hacer con BodyParser
	//if err := json.Unmarshal(c.Body(),&u); err != nil{return err;}
	

	type response struct{
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil{return err;}
	//sanitize this
	//res.user
	//res.pass

	id := CheckCredentials(res.User, res.Pass);

	if( id != -1){ // se ha encontrado

		//mover este bloque a una funcion de jwt
		token := jwt.CreateToken(id);

		cookie := new(fiber.Cookie);
		cookie.Name = "token";
		cookie.Value = token;
		//cookie.Expires = time.Now().Add(time.Minute*10);
		cookie.Expires = time.Now().Add(time.Minute*1);
		cookie.HTTPOnly = true;
		//cookie.Secure = false; //por ahora

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
