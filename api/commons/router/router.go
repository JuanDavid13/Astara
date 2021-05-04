package router;

import (

	"github.com/gofiber/fiber/v2"

	 . "astara/controllers"
	 . "astara/commons/jwt"
)


//Router declarations
func RouterSetUp(app *fiber.App){
	api := app.Group("/api/v1");

	login := api.Group("/login");
		login.Post("/", CheckUser);
		login.Post("/check", Check);
		login.Get("/validate", Validate);
}
