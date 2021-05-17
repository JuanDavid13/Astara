package controllers

import (
	"fmt"

	"encoding/json"
	str "strings"

	"github.com/gofiber/fiber/v2"

	cookie "astara/commons/cookie"
	jwt "astara/commons/jwt"
	. "astara/models"
)

func GetAreas(c *fiber.Ctx) error {
	fmt.Println("get area");
	cl := c.Locals("claims").(jwt.Claims);

	//user := jwt.GetUser(c);
	areas := GetAreasById(cl.User,cl.Rol);

	Areas,err := json.Marshal(areas);
	if err != nil{ panic(err); }

	c.Status(200);
	return c.JSON(string(Areas));
}

//We can do a lower case transform, before or after we process the string
//but I'm going to assume that it is always goint to be lower case
func NameToSlug(name string) string { return str.ReplaceAll(name," ","-"); }
func SlugToName(slug string) string { return str.ReplaceAll(slug,"-"," "); }

func AreaCheck(c *fiber.Ctx) error {

	type name struct { Name string `json:"name"`; }
	n := name{};

	if err := json.Unmarshal(c.Body(),&n); err != nil{return err;}

	//if c.Cookies("token") == "" { return c.SendStatus(401); }
	if cookie.CheckIsEmpty(c.Cookies("token")) {
		//return c.JSON(fiber.Map{
		//	"correspond":false,
		//})
		return c.SendStatus(401);
	}
	cl := c.Locals("claims").(jwt.Claims);
	if targets, found := CheckUserArea(cl.User,n.Name,cl.Rol); found {
		c.Status(200);
		return c.JSON(fiber.Map{
			"correspond":true,
			"targets":targets,
		}); 
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{
			"correspond":false,
		}); 
	}
	//return c.SendStatus(400);
}
