package main;

import (
	//fiber http server
	"github.com/gofiber/fiber/v2"

	//cors
	//"github.com/gofiber/fiber/v2/middleware/cors"

	//Json encoding
	//"enconding/json"

	"astara/commons"

)

func main(){
	//router := commons.Router{};
	db := commons.Db{};
	db.Open("root","")


	app := fiber.New(fiber.Config{
//		Prefork:				true,
	});

	app.Get("/api/v1/", func(c *fiber.Ctx) error{
		return c.SendString("Hello world");
	});

	app.Listen(":3000");
}
