package main

import (
	//fiber http server
	"github.com/gofiber/fiber/v2"
	//cors
	"github.com/gofiber/fiber/v2/middleware/cors"

	//Json encoding
	_ "encoding/json"

	C "astara/commons"
	M "astara/models"
)

func main(){
	//router := commons.Router{};

	db := C.Db{};
	db.New();
	datab := db.Open("root","");
	
	t := M.Target{};
	t.GetAllTargets(datab);

	app := fiber.New(fiber.Config{
//		Prefork:				true,
	});

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081",
	}));

	app.Get("/api/v1/", func(c *fiber.Ctx) error{
		return c.JSON(t.GetAllTargets(datab));
		//return c.SendString("Hello world");
	});

	app.Listen(":3000");
}
