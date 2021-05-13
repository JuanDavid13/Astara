package controllers

import (
	"fmt"

	"encoding/json"
	str "strings"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	. "astara/models"
)

func GetAreas(c *fiber.Ctx) error {

	user := jwt.GetUser(c.Cookies("token"))
	areas := GetAreasById(user);

	Areas,err := json.Marshal(areas);
	if err != nil{ panic(err); }

	for i:=0; i< len(areas); i++ {
		fmt.Println(areas[i].Name);
		fmt.Println(NameToSlug(areas[i].Name));
	}

	c.Status(fiber.StatusOK);
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

	fmt.Println(n);
	fmt.Println(n.Name);

	//a cookie controler needs to be done
	if c.Cookies("token") == ""{ return c.SendStatus(401); }
	user := jwt.GetUser(c.Cookies("token"))

	targets, found := CheckUserArea(user,n.Name)

	if  found { 
		c.Status(200);
		return c.JSON(fiber.Map{
			"correspond":true,
			"targets":targets,
		}); 
	}

	c.Status(200);
	return c.JSON(fiber.Map{
		"correspond":false,
	}); 
}
