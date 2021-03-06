package router

import (
	//"fmt"

	"github.com/gofiber/fiber/v2"

	. "astara/commons/jwt"
	. "astara/controllers"
)

func middleware(c *fiber.Ctx) error {
	if Validate(c) { return c.Next(); }
	return c.SendStatus(401);
}

//Router declarations
func RouterSetUp(app *fiber.App) {
	api := app.Group("/api/v1");

	login := api.Group("/login");
		login.Post("/", CheckUser);
		login.Post("/check", Check);
		login.Post("/create", CreateUser);

	auth := api.Group("/auth");
		auth.Get("/validate",AuthValidate);
		auth.Get("/logout",LogOut);

	user:= api.Group("/user");
		user.Use(middleware);

		task := user.Group("/task");
			task.Post("/create",CreateTask);
			task.Post("/delete",DeleteTask);
			task.Post("/edit",EditTask);
			task.Post("/goal-tasks",GetTasksOfGoal);
			task.Post("/check",CheckTask);

	goal := api.Group("/goal");
		goal.Use(middleware);
		goal.Post("/create",CreateGoal);
		goal.Post("/delete",DeleteGoal);
		goal.Post("/edit",EditGoal);

		info := user.Group("/profile");
			info.Get("/info",GetInfo);
			info.Post("/update",UpdateUser);
			info.Post("/checkpass",CheckPass);
			info.Post("/changePass",UpdatePwd);

	area := api.Group("/area");
		area.Use(middleware);
		area.Get("/",GetAreas);
		area.Post("/correspond",AreaCheck);
		area.Post("/create",CreateArea);
		area.Post("/delete",DeleteArea);
		area.Post("/remove-target",RemoveTarget);
		area.Post("/edit",ChangeAreaName);
		area.Get("/main/goals",GetMainGoals);
		area.Get("/main/tasks",GetMainTasks);

		area.Get("/:slug/paginated-tasks/:size/:paginated",GetPagTasks);
		area.Get("/:slug/paginated-goals/:size/:paginated",GetPagGoals);
}

