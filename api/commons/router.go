package commons;

import (
	"github.com/gofiber/fiber/v2"
)

//type Router struct{
//	Nombre string;
//}

func Llega(c *fiber.Ctx) error{
	return c.SendString("llega también");
}

//Router declarations
func RouterSetUp(app *fiber.App){
	api := app.Group("/api/v1");
	login := api.Group("/login");

	login.Get("/", Llega)



	app.Get("/api/v1/prueba", func(c *fiber.Ctx) error {
		return c.SendString("fundiona");
	})

}
