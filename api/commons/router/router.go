package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	. "astara/commons/jwt"
	. "astara/controllers"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("middleware");
	if !Validate(c) { return c.SendStatus(/*401*/200);	}
		return c.Next();
}

//Router declarations
func RouterSetUp(app *fiber.App){
	api := app.Group("/api/v1");

	login := api.Group("/login");
		login.Post("/", CheckUser);
		login.Post("/check", Check);

	auth := api.Group("/auth");
		auth.Get("/validate",AuthValidate);

	user:= api.Group("/user");
		user.Use(middleware);
		user.Get("/targets",GetTargets);

	area := api.Group("/area");
		area.Use(middleware);
		area.Get("/",GetAreas);
		area.Post("/correspond",AreaCheck);
}

