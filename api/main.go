package main

import (
	"os"

	//fiber http server
	"github.com/gofiber/fiber/v2"
	//cors
	"github.com/gofiber/fiber/v2/middleware/cors"

	//env access
	env "github.com/joho/godotenv"

	. "astara/commons/router"
)

func main(){

	err := env.Load();
	if err != nil {panic(err);}

	app := fiber.New(fiber.Config{
		Prefork:				true,
	});

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONT_BASE_URL"),
		AllowCredentials: true,
	}));

	RouterSetUp(app);

	app.Listen(":3000");
}
