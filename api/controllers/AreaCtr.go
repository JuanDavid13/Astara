package controllers

import (
	"fmt"

	"encoding/json"
	str "strings"

	"github.com/gofiber/fiber/v2"

	//cookie "astara/commons/cookie"
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

	type slug struct { Slug string `json:"slug"`; }
	s := slug{};

	if err := json.Unmarshal(c.Body(),&s); err != nil{ return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);

	c.Status(200);
	areaId, name := CheckUserArea(cl.User,s.Slug,cl.Rol);
	if areaId == -1 || areaId == 0 {
		c.Status(200);
		return c.JSON(fiber.Map{
			"correspond":false,
		}); 
	}
	
	deleteable := AreaIsDeleteable(areaId, cl.Rol);
	if !deleteable{ 
		return c.JSON(fiber.Map{
			"correspond":true,
			"deleteable":false,
			"areaName": name,
		}); 
	}

	return c.JSON(fiber.Map{
		"correspond":true,
		"deleteable":true,
		"areaName": name,
	}); 
}

func CreateArea(c *fiber.Ctx) error {

	type response struct{ Name string `json:"name"`; }
	res := response{};

	if err := json.Unmarshal(c.Body(), &res); err != nil {
		c.Status(400);
		return c.JSON(fiber.Map{
			"added":false,
		});
	}

	cl := c.Locals("claims").(jwt.Claims);

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


func DeleteArea(c *fiber.Ctx) error {

	type slug struct { Slug string `json:"slug"`; }
	slugreq := slug{}; //slugreq -> slug request


	if err := json.Unmarshal(c.Body(),&slugreq); err != nil { panic(err); }

	if slugreq.Slug == "" { return c.SendStatus(400);}

	cl := c.Locals("claims").(jwt.Claims);

	c.Status(200);
	if deleted := DelArea(cl.User,cl.Rol,slugreq.Slug); !deleted {
		return c.JSON(fiber.Map{
			"deleted": false,
		});
	}
	return c.JSON(fiber.Map{
		"deleted": true,
	});
}

func ChangeAreaName(c *fiber.Ctx) error {

	type response struct { 
		Area string `json:"area"`; 
		Name string `json:"name"`; 
	}
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil{ return c.SendStatus(400); }

	cl:= c.Locals("claims").(jwt.Claims);

	slug := NameToSlug(res.Name); //can obviate this
	c.Status(200);
	if !ChangeName(cl.User, cl.Rol, res.Area, res.Name, slug){
					return c.JSON(fiber.Map{ "changed":false, });
	}else{	return c.JSON(fiber.Map{ "changed":true, }); }
} 

func RemoveTarget(c *fiber.Ctx) error {
	type response struct { Id int `json:"id"`; }
	res := response{};

	fmt.Println("Id");
	fmt.Println(res.Id);

	err := json.Unmarshal(c.Body(), &res);
	if err != nil { panic(err); /*return c.SendStatus(400);*/}

	if res.Id <= 0 { return c.SendStatus(400); }

	cl := c.Locals("claims").(jwt.Claims);

	if !RmvTarget(cl.User, res.Id, cl.Rol){
		return c.JSON(fiber.Map{
			"error":true,
		});
	}else{
		return c.JSON(fiber.Map{
			"error":false,
			"deleted":true,
		});
	}
}
