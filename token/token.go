package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetLoggedIn(c echo.Context) map[string]interface{} {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	email := claims["email"].(string)
	uid := claims["uid"].(string)

	login := make(map[string]interface{})
	login["name"] = name
	login["email"] = email
	login["uid"] = uid

	fmt.Println(claims)
	return login
}
