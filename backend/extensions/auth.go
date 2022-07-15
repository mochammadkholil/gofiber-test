package extensions

import (
	"backendsvc/models"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/valyala/fasthttp"
)

func ExtractToken(r *fasthttp.Request) string {
	bearToken := string(r.Header.Peek("Authorization"))
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *fasthttp.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func Auth(excluded []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessable := false
		var err error
		for _, i := range excluded {
			if strings.Contains(c.Path(), i) {
				accessable = true
			}
		}

		if !accessable {
			token, err := VerifyToken(c.Request())
			if err == nil {
				c.Locals("user", token)
				c.Next()
			} else {
				c.JSON(models.ApiResponse{
					Success: false,
					Message: fmt.Sprintf("don't have permission, %v", err.Error()),
				})
			}
		} else {
			c.Next()
		}

		return err
	}
}
