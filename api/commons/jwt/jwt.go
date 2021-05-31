package jwt 

import (
	"fmt"

	"os"
	"time"
	str "strconv"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"

	"astara/commons/cookie"
)

type Claims struct{
	Authorized bool `json:"authorized"`;
	Exp int64 `json:"exp"`;
	User int `json:"user"`;
	Rol string `json:"rol"`;
	jwt.StandardClaims
}

func newClaim(user int, rol string) Claims {
	dur, _ := str.Atoi(os.Getenv("CK_DUR"));
	return Claims{
		true,
		time.Now().Add(time.Duration(dur)).Unix(),
		user,
		rol,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(dur)).Unix(),
		},
	}
}

func CreateToken(user int, rol string) string {
	claims := newClaim(user, rol);
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims);

	tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
	if err != nil {panic(err);}

	return tokenString;
}

func ParseToken(tokenString string) (*Claims, bool){
	fmt.Println("tokenString");
	fmt.Println(tokenString);

	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	});
	if err != nil{ return nil, true; }

	claims ,ok := token.Claims.(*Claims);
	if !ok {panic(ok)}

	return claims,true;
}

func IsEmpty(cookieString string) bool {
	if cookieString != "" { return false; }
	return true;
}

//resource validation
func Validate(c *fiber.Ctx) bool {
	fmt.Println("validate");
	//fmt.Println(string(c.Request().Header.Peek("Origin")));

	//isValid := false;
	if IsEmpty(c.Cookies("token")){
		return false;
	}else{
		if claims, valid := ParseToken(c.Cookies("token")); !valid {
			return false;
		}else{
			cl := Claims(*claims);
			c.Locals("claims", cl);

			//renew the token
			newToken := CreateToken(cl.User, cl.Rol);
			newCookie := cookie.CreateCookie(newToken);
			c.Cookie(newCookie);

			return true;
		}
	}
	//if !IsEmpty(c.Cookies("token")){
		//if claims, valid := ParseToken(c.Cookies("token")); valid {
		//	cl := Claims(*claims);
		//	c.Locals("claims", cl);

		//	//renew the token
		//	newToken := CreateToken(cl.User, cl.Rol);
		//	newCookie := cookie.CreateCookie(newToken);
		//	c.Cookie(newCookie);

		//	isValid = true;
	//	}
	//}
}

//page validation
func AuthValidate(c *fiber.Ctx) error {
	fmt.Println("Auth validation");

	if !IsEmpty(c.Cookies("token")){
		if _, valid := ParseToken(c.Cookies("token")); valid { 
			return c.SendStatus(200); 
		}
	}
	fmt.Println(IsEmpty(c.Cookies("token")));
	return c.SendStatus(401);
}

//delete from this
func GetUser(c *fiber.Ctx) int { 
	fmt.Println("trying to get the user");
	//fmt.Println(reflect.TypeOf(token));
	//claims := token.Claims.(*Claims);
	//return claims.User; 
	return 0;
}

//re-enfocar
//delete from this
func GetExp(claims Claims) int { return int(claims.Exp); }

func RenewExp(claims *Claims) string{
	//claims := token.Claims.(*Claims);

	//dur, _ := str.Atoi(os.Getenv("CK_DUR"));

	//claims.Exp = time.Now().Add(time.Duration(dur)).Unix();
	//claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(dur)).Unix(); //me lo puedo saltar si quiero

	//
	//tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
	//if err != nil { panic(err); }

	return "tokenString";
}
