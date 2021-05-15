package jwt 

import (
	"fmt"
	"reflect"

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

var expTime = time.Minute*10;

func newClaim(user int) Claims {
	return Claims{
		true,
		time.Now().Add(expTime).Unix(),
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expTime).Unix(),
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

func ParseToken(tokenString string) (*Claims, bool){
	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	});
	if err != nil{ 
		//panic(err);
		return nil, true;
	}

	claims ,ok := token.Claims.(*Claims);
	if !ok {panic(ok)}

	return claims,true;
}

/*func GetTokenClaims(tokenString string) *Claims{

	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	});
	if err != nil{panic(err);}

	claims , ok := token.Claims.(*Claims);
	if !ok {panic(ok)}

	return claims;
}*/

func IsEmpty(cookieString string) bool {
	fmt.Println(cookieString);
	if cookieString != "" { return false; }
	return true;
}

func Validate(c *fiber.Ctx) bool {
	fmt.Println("validate");
	//fmt.Println(string(c.Request().Header.Peek("Origin")));

	valid := false;
	if !IsEmpty(c.Cookies("token")){
		claims, valid := ParseToken(c.Cookies("token"));
		fmt.Printf("%+v\n",claims);
		if valid { c.Locals("claims", claims); }
	}
	return valid;
}

func AuthValidate(c *fiber.Ctx) error {
	fmt.Println("Auth validation");

	if !IsEmpty(c.Cookies("token")){
		fmt.Println("no está vacio");
		if _, valid := ParseToken(c.Cookies("token")); valid { return c.SendStatus(204); }
	}
	return c.SendStatus(401);
}

func GetUser(c *fiber.Ctx) int { 
	fmt.Println("trying to get the user");
	token := c.Locals("token");
	fmt.Println(reflect.TypeOf(token));
	//claims := token.Claims.(*Claims);
	//return claims.User; 
	return 0;
}
//re-enfocar
func GetExp(claims Claims) int { return int(claims.Exp); }

func RenewExp(token *jwt.Token) string{
	claims := token.Claims.(*Claims);

	claims.Exp = time.Now().Add(expTime).Unix();
	claims.StandardClaims.ExpiresAt = time.Now().Add(expTime).Unix(); //me lo puedo saltar si quiero

	tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
	if err != nil { panic(err); }

	return tokenString;
}
