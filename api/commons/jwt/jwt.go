package jwt 

import (
	"os"
	"time"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(user int) string {
		token := jwt.New(jwt.SigningMethodHS512);
		claims := token.Claims.(jwt.MapClaims);
		claims["authorized"] = true;
		claims["user"] = user;
		claims["exp"] = time.Now().Add(time.Minute*10).Unix();

		tokenString, err := token.SignedString([]byte(os.Getenv("SCRT")));
		if err != nil {panic(err);}

		return tokenString;
}

func CheckToken(tokenString string) /*jwt.MapClaims*/ bool {

	claims := jwt.MapClaims{};
	_ , err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SCRT")),nil
	})

	if err != nil{
		fmt.Println("El token no es correcto");
		return false;
	}else{
		fmt.Println("El token es correcto");
		return true;
	}
}

func CheckExpTime(claims jwt.MapClaims) bool {
	if int64(claims["exp"].(float64)) <= time.Now().Unix() {return false;}
	return true;

}

