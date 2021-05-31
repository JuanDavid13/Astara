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

func CreateArea(c *fiber.Ctx) error {
	cl := c.Locals("claims").(jwt.Claims);

	type response struct{ Name string `json:"name"`; }
	res := response{};

	if err := json.Unmarshal(c.Body(), &res); err != nil {
		c.Status(400);
		return c.JSON(fiber.Map{
			"added":false,
		});
	}

	if res.Name == "" {
		c.Status(400);
		return c.JSON(fiber.Map{
			"added":false,
		});
	}

	slug := NameToSlug(res.Name);

	created := CreateNewArea(cl.User,cl.Rol, res.Name, slug );
	if !created {
		c.Status(200);
		return c.JSON(fiber.Map{
			"added":false,
		});
	}else{
		areas := GetAreasById(cl.User,cl.Rol);
		Areas,err := json.Marshal(areas);
		if err != nil{ panic(err); }

		c.Status(200);
		return c.JSON(fiber.Map{
			"added":true,
			"areas":string(Areas),
		});
	}
}
