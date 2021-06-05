package controllers

import (
	"fmt"
	//"reflect"

	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	jwt "astara/commons/jwt"
	"astara/commons/cookie"
	. "astara/models"
)

//sanitize
func CheckUser(c *fiber.Ctx) error {
	fmt.Println("Check user");
	u := User{};
	if err := json.Unmarshal(c.Body(),&u); err != nil{ return c.SendStatus(400); }

	c.Status(200);
	if IsRegistered(u.User){ return c.JSON(fiber.Map{"found":true,})
	}else{ return c.JSON(fiber.Map{"found":false,})}
}

func logUser(c *fiber.Ctx, name string, id int, rol string) {
	token := jwt.CreateToken(id,rol); // Create a token
	cookie := cookie.CreateCookie(token); // Create a cookie with that token value
	c.Cookie(cookie); // Set the cookie
}

func Check(c *fiber.Ctx) error {
	fmt.Println("Check");
	//other way to do it
	type response struct{
		User string `json:"user"`
		Pass string `json:"pass"`
	};
	res := response{};

	if err := json.Unmarshal(c.Body(),&res); err != nil{/*return err;*/ c.SendStatus(400);}

	c.Status(200);
	if id, correct := CheckCredentials(res.User, res.Pass); !correct {
		return c.JSON(fiber.Map{"logged":"false"});
	}else{
		name, rol := GetBasicUserInfo(id);
		if name == "" || rol == "" { return c.SendStatus(400); }
		logUser(c, name, id, rol);
		return c.JSON(fiber.Map{"logged":"true"});
	}
}

func LogOut(c *fiber.Ctx) error {
	if !cookie.CheckIsEmpty(c.Cookies("token")) {
		cookie := cookie.CreateExpiredCookie();
		c.Cookie(cookie);
		return c.SendStatus(200);
	}
	return c.SendStatus(400);
}

func CreateUser(c *fiber.Ctx) error {
	type response struct {
		User string `json:"user"`;
		Pass string `json:"pass"`;
		Email string `json:"email"`;
	}
	res := response{};
	if err := json.Unmarshal(c.Body(),&res); err != nil {
		panic(err);
	}

	if IsRegisteredWithEmail(res.User, res.Email) {
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": true,
			"alreadyCreated": true,
			"created": false,
		});
	}
	
	id, created := CreateNewUser(res.User,res.Pass, res.Email);
	if !created {
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": true,
			"alreadyCreated": false,
			"created": false,
		});
	}else{
		CreateIndexArea(id);
		logUser(c, res.User, id, "user");
		c.Status(200);
		return c.JSON(fiber.Map{
			"error": false,
			"alreadyCreated": false,
			"created": true,
		});

	}
}

func GetInfo(c *fiber.Ctx) error {
	cl := c.Locals("claims").(jwt.Claims);
	if name, email, theme, err := GetBasicInfo(cl.User,cl.Rol); err != false {
		panic(err);
		//return c.SendStatus(200);
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{
			"name":name,
			"email":email,
			"theme":theme,
		})
	}
}

func CheckPass(c *fiber.Ctx) error {
	cl := c.Locals("claims").(jwt.Claims);

	type passRes struct{ Pass string `json:"password"`; }
	pass := passRes{};

	if err := json.Unmarshal(c.Body(),&pass); err != nil { panic(err); }

	c.Status(200);
	same, err := ComparePass(cl.User,cl.Rol,pass.Pass);
	if err { 
		return c.JSON(fiber.Map{
			"same":false,
		})
	}
	if !same{
		return c.JSON(fiber.Map{
			"same":false,
		})
	}else{
		return c.JSON(fiber.Map{
			"same":true,
		})
	}
}

func UpdatePwd(c *fiber.Ctx) error {
	type response struct { Pass string `json:"pass"`; }
	res := response{}

	if err := json.Unmarshal(c.Body(),&res); err != nil { panic(err); }

	cl := c.Locals("claims").(jwt.Claims);

	c.Status(200);
	if !UpdateUserPwd(cl.User,cl.Rol,res.Pass) {
		return c.JSON(fiber.Map{ "updated":false, });
	}else{
		return c.JSON(fiber.Map{ "updated":true, });
	}

	return c.SendStatus(200);
}

func UpdateUser(c *fiber.Ctx) error {
	cl := c.Locals("claims").(jwt.Claims);

	type response struct {
		Username *string `json:"username"`;
		Email *string `json:"email"`;
		Theme *bool `json:"theme"`;
	}

	type changes struct{
			Changes response `json:"changes"`;
	}
	//res := response{};
	res:= changes{};

	if err := json.Unmarshal(c.Body(), &res); err != nil { return c.SendStatus(400) /*panic(err);*/ }

	query :=createUpdateQuery(cl.User, res.Changes.Username, res.Changes.Email, res.Changes.Theme);
	if query == nil{ return c.SendStatus(400); }


	var changesArr []string;
	if res.Changes.Username != nil {
		if UserTakenbyName(*res.Changes.Username){
			c.Status(200);
			return c.JSON(fiber.Map{
				"updated":false, 
			});
		}
		changesArr = append(changesArr, *res.Changes.Username); }
	if res.Changes.Email != nil {
		if UserTakenbyEmail(*res.Changes.Username){
			c.Status(200);
			return c.JSON(fiber.Map{
				"updated":false, 
			});
		}
		changesArr = append(changesArr, *res.Changes.Email); }
	if res.Changes.Theme != nil { 
		if *res.Changes.Theme == true {
						changesArr = append(changesArr,"1"); 
		}else{	changesArr = append(changesArr,"0"); }
		//strconv.FormatBool(*res.Changes.Theme)
	}

	if !UpdateUserInfo(cl.User, cl.Rol, *query, changesArr) {
		c.Status(200);
		return c.JSON(fiber.Map{ "updated":false, });
	}else{
		c.Status(200);
		return c.JSON(fiber.Map{ "updated":true, });
	}
}

func createUpdateQuery(User int, username, email *string, theme *bool) *string {
	if username == nil && email == nil && theme == nil { return nil; }

	query := "UPDATE `Users` SET ";

	if username != nil { // cambia
			query += "`Name` = ?"
	}

	if email != nil { // cambia
		if username != nil { query += ", `Email` = ?";
		}else{ query += "`Email` = ?"; }
	}

	if theme != nil { // cambia
		if username != nil || email != nil { query += ", `Theme` = ?";
		}else{ query += "`Theme` = ?"; }
	}

	query += " WHERE `Id` = " + strconv.Itoa(User) + ";";

	return &query;
}
