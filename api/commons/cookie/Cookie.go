package cookie

import (
	str "strconv"
	"time"
	"os"

	//jwt "github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
)

func CreateCookie(tokenString string) *fiber.Cookie {
	cookie := new(fiber.Cookie);
	cookie.Name = "token";
	cookie.Value = tokenString;
	dur,err := str.Atoi(os.Getenv("CK_DUR"));
	if err != nil { panic(err); }

	cookie.Expires = time.Now().Add(time.Minute*(time.Duration(dur)));

	cookie.HTTPOnly = true;
	//cookie.Secure = true;

	return cookie;
}

func CreateExpiredCookie() *fiber.Cookie {
	cookie := new(fiber.Cookie);
	cookie.Name = "token";
	cookie.Value = "";

	t, err := time.Parse(time.RFC3339,"2009-11-10T23:00:00.000Z");
	if err != nil { panic(err); }
	cookie.Expires = t;

	cookie.HTTPOnly = true;
	//cookie.Secure = true;

	return cookie;
}

func CheckIsEmpty(cookieString string) bool { return cookieString == ""; }
