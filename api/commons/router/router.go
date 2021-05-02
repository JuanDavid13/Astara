package router;

import (
	"github.com/gofiber/fiber/v2"

	 . "astara/controllers"
)

//type Router struct{
//	Nombre string;
//}

//handler example
//func Llega(c *fiber.Ctx) error{
//	return c.SendString("llega también");
//}

//Router declarations
func RouterSetUp(app *fiber.App){
	api := app.Group("/api/v1");

	login := api.Group("/login");
		login.Post("/", IsRegistered);
		login.Post("/check", CheckCredentials);

	auth := api.Group("/auth");
		auth.Post("/", CheckToken);




	app.Get("/api/v1/prueba", func(c *fiber.Ctx) error {
		return c.SendString("fundiona");
	})

}
