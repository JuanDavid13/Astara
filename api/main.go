package main

import (
	"fmt"
	//"go/token"
	"os"
	"time"
	"reflect"

	//fiber http server
	"github.com/gofiber/fiber/v2"
	//cors
	"github.com/gofiber/fiber/v2/middleware/cors"
	//session
	"github.com/gofiber/fiber/v2/middleware/session"

	//env access
	env "github.com/joho/godotenv"

	//jwt
	//jwt "github.com/dgrijalva/jwt-go"

	//Json encoding
	_ "encoding/json"

	C "astara/commons"
	M "astara/models"
)

func main(){
	//router := commons.Router{};

	err := env.Load();
	if err != nil {panic(err);}

	db := C.Db{};
	db.New();
	datab := db.Open("root","");
	
	t := M.Target{};
	t.GetAllTargets(datab);

	app := fiber.New(fiber.Config{
//		Prefork:				true,
	});

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONT_BASE_URL"),
		AllowCredentials: true,
	}));

	store := session.New();

	C.RouterSetUp(app);

	app.Get("/api/v1/", func(c *fiber.Ctx) error{

		tokenString := C.CreateToken();
		claims := C.CheckToken(tokenString);
		exp := claims["exp"];
		fmt.Println(exp);
		//fmt.Println(reflect.TypeOf(exp));
		//fmt.Println(int64(exp.(float64)));
		fmt.Println(C.CheckExpTime(claims))
		fmt.Println(claims);
		fmt.Println(reflect.TypeOf(C.CheckToken(tokenString)));

		//token := jwt.New(jwt.SigningMethodHS512);
		//claims := token.Claims.(jwt.MapClaims);

		//claims["authorized"] = true;
		//claims["user"] = "Ronaldo";
		//claims["exp"] = time.Now().Add(time.Minute*10).Unix();

		//tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
		//if err != nil {panic(err);}

		//fmt.Println(tokenString);


		//fmt.Println(c.Request().Header)
		//fmt.Println(c.Get("Authorization"))

		cookie := new(fiber.Cookie);
		cookie.Name = "token";
		cookie.Value = "alguno";
		cookie.Expires = time.Now().Add(1 * time.Hour);

		c.Cookie(cookie);

		sess, err := store.Get(c);
		if err != nil{panic(err);}
		defer sess.Save()

		//name := sess.Get("name");
		//fmt.Println(name);

		//sess.Set("name","John");
		//sess.Set("rol","1");

		//fmt.Print(sess.Get("name"));

		
		//name = sess.Get("name");
		//fmt.Println(name);

		//sess.Destroy();

		return c.JSON(t.GetAllTargets(datab));
		//return c.SendString("Hello world");
	});

	app.Listen(":3000");
}
