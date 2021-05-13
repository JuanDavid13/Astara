package jwt 

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
)

type Claims struct{
	Authorized bool `json:"authorized"`;
	Exp int64 `json:"exp"`;
	User int `json:"user"`;
	jwt.StandardClaims
}

func newClaim(user int) Claims {
	return Claims{
		true,
		time.Now().Add(time.Minute*10).Unix(),
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*10).Unix(),
		},
	}
}

func CreateToken(user int) string {
		
	claims := newClaim(user);
		
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims);

	tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
	if err != nil {panic(err);}

	return tokenString;
}

func CheckToken(tokenString string) (bool) {

	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	});
	if err != nil{panic(err);}

	_ ,ok := token.Claims.(*Claims);
	if !ok {panic(ok)}


	return true;
}

func CheckExpTime(exp int64) bool {
	if exp <= time.Now().Unix() {return false;}
	return true;
}

func IsEmpty(cookie string) bool {
	if cookie != "" { return false; }
	return true;
}

func Validate(c *fiber.Ctx) error{

	if !IsEmpty(c.Cookies("token")){
		if CheckToken(c.Cookies("token")) {
			c.Status(200);
			return c.JSON(fiber.Map{"valid":"true"});
		}
	}
	c.Status(200);
	return c.JSON(fiber.Map{"valid":"false"});
}

func GetUser(tokenString string) int {

	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	});
	if err != nil{panic(err);}

	claims ,ok := token.Claims.(*Claims);
	if !ok {panic(ok)}

	return claims.User;
}
