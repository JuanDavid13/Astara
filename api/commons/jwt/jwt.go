package jwt 

import (
	"fmt"
	//"reflect"

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
	jwt.StandardClaims
}

func newClaim(user int) Claims {
	dur, _ := str.Atoi(os.Getenv("CK_DUR"));
	return Claims{
		true,
		time.Now().Add(time.Duration(dur)).Unix(),
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(dur)).Unix(),
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
	fmt.Println("IsEmpty");
	if cookieString != "" { 
		fmt.Println(false);
		fmt.Println();
		return false; 
	}
	fmt.Println(false);
	fmt.Println();
	return true;
}

//resource validation
func Validate(c *fiber.Ctx) bool {
	fmt.Println("validate");
	//fmt.Println(string(c.Request().Header.Peek("Origin")));

	valid := false;
	if !IsEmpty(c.Cookies("token")){
		claims, valid := ParseToken(c.Cookies("token"));
		cl := Claims(*claims);
		//fmt.Printf("%+v\n",claims);
		if valid { 
			c.Locals("claims", cl);

			newToken := CreateToken(cl.User);
			newCookie := cookie.CreateCookie(newToken);
			c.Cookie(newCookie);

			valid = true;
		}
		fmt.Println(valid);
		fmt.Println();
		return valid;
	}
	fmt.Println(valid);
	fmt.Println();
	return valid;
}

//page validation
func AuthValidate(c *fiber.Ctx) error {
	fmt.Println("Auth validation");

	if !IsEmpty(c.Cookies("token")){
		if _, valid := ParseToken(c.Cookies("token")); valid { 
			fmt.Println(200);
			fmt.Println();
			return c.SendStatus(200); 
		}
	}

	fmt.Println(401);
	fmt.Println();
	return c.SendStatus(401);
}

func GetUser(c *fiber.Ctx) int { 
	fmt.Println("trying to get the user");
	//fmt.Println(reflect.TypeOf(token));
	//claims := token.Claims.(*Claims);
	//return claims.User; 
	return 0;
}

//re-enfocar
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
