package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rover10/mydocapp/server"
)

type Handler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *Handler) Private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	email := claims["email"].(string)
	//sid := claims["sid"].(string)

	fmt.Println(claims)
	return c.String(http.StatusOK, "Welcome "+name+"!"+email)
}

// This is the api to refresh tokens
// Most of the code is taken from the jwt-go package's sample codes
// https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func (h *Handler) Token(c echo.Context) error {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		//serv := server.NewServer(config.Config{})
		return []byte(server.SECRET_PASSWORD), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["sub"].(float64)) == 1 {
			name, ok1 := claims["name"].(string)
			uid, ok2 := claims["uid"].(string)
			email, ok3 := claims["email"].(string)
			if ok1 && ok2 && ok3 {
				newTokenPair, err := GenerateTokenPair(name, uid, email)
				if err != nil {
					return err
				}
				return c.JSON(http.StatusOK, newTokenPair)
			}
		}

		return echo.ErrUnauthorized
	}

	return err
}
