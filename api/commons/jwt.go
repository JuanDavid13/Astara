package commons

import (
	"os"
	"time"
//	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken() string {

		token := jwt.New(jwt.SigningMethodHS512);
		
		claims := token.Claims.(jwt.MapClaims);

		claims["authorized"] = true;
		claims["user"] = "Ronaldo";
		claims["exp"] = time.Now().Add(time.Minute*10).Unix();

		tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
		if err != nil {panic(err);}

		return tokenString;
}

func CheckToken(tokenString string) jwt.MapClaims{

	claims := jwt.MapClaims{};
	_ , _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	})

	return claims;
}

func CheckExpTime(claims jwt.MapClaims) bool {
	if int64(claims["exp"].(float64)) <= time.Now().Unix() {return false;}
	return true;

}
