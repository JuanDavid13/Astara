package router

import (
	"github.com/gofiber/fiber/v2"

	. "astara/commons/jwt"
	. "astara/controllers"
)

//Router declarations
func RouterSetUp(app *fiber.App){
	api := app.Group("/api/v1");

	login := api.Group("/login");
		login.Post("/", CheckUser);
		login.Post("/check", Check);
		login.Get("/validate", Validate);
		//login.Get("/parse", jwt.CheckToken);

	user:= api.Group("/user");
		user.Get("/areas",GetAreas);
		user.Get("/targets",GetTargets);
}
