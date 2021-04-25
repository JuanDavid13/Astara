package main

import (
	"fmt"
	"os"

	//fiber http server
	"github.com/gofiber/fiber/v2"
	//cors
	"github.com/gofiber/fiber/v2/middleware/cors"
	//session
	"github.com/gofiber/fiber/v2/middleware/session"

	//env access
	"github.com/joho/godotenv"

	//Json encoding
	_ "encoding/json"

	C "astara/commons"
	M "astara/models"
)

func main(){
	//router := commons.Router{};

	err := godotenv.Load();
	if err != nil {panic(err);}
	fmt.Println(os.Getenv("SCRT"));

	db := C.Db{};
	db.New();
	datab := db.Open("root","");
	
	t := M.Target{};
	t.GetAllTargets(datab);

	app := fiber.New(fiber.Config{
//		Prefork:				true,
	});

	app.Use(cors.New(cors.Config{
		//AllowOrigins: "http://localhost:8081",
		AllowOrigins: "http://localhost:8081",
		AllowCredentials: true,
	}));

	store := session.New();

	app.Get("/api/v1/", func(c *fiber.Ctx) error{

		sess, err := store.Get(c);
		if err != nil{panic(err);}
		defer sess.Save()

		//name := sess.Get("name");
		//fmt.Println(name);

		//sess.Set("name","John");
		//sess.Set("rol","1");

		fmt.Print(sess.Get("name"));

		
		//name = sess.Get("name");
		//fmt.Println(name);

		//sess.Destroy();

		return c.JSON(t.GetAllTargets(datab));
		//return c.SendString("Hello world");
	});

	app.Listen(":3000");
}
