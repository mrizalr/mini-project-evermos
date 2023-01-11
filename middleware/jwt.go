package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/mini-project-evermos/model"
	"github.com/spf13/viper"
)

func validateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token isn't valid : unexpected signing method %v: ", t.Method)
		}
		return []byte(viper.GetString("secret_key")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token isn't valid")
	}

	return &claims, nil
}

func Auth(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Unauthorized",
			Errors:  []string{"you are not authorized to access this resource. Please authenticate or contact the administrator for assistance"},
			Data:    nil,
		})
	}

	headers := strings.Split(authorizationHeader, " ")
	if len(headers) < 2 || headers[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Invalid header format",
			Errors:  []string{"Invalid header format: must use `Bearer (token)` format"},
			Data:    nil,
		})
	}

	tokenString := headers[1]

	mapClaims, err := validateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Unauthorized",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	expStr, ok := (*mapClaims)["exp"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Invalid token",
			Errors:  []string{"your token isn't valid"},
			Data:    nil,
		})
	}

	exp, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Invalid token",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	if time.Now().Unix() > exp {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Expired token",
			Errors:  []string{"you are not authorized to access this resource. Please authenticate or contact the administrator for assistance"},
			Data:    nil,
		})
	}

	userID, ok := (*mapClaims)["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Invalid token",
			Errors:  []string{"your token isn't valid"},
			Data:    nil,
		})
	}

	userRole, ok := (*mapClaims)["user_role"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status:  false,
			Message: "Invalid token",
			Errors:  []string{"your token isn't valid"},
			Data:    nil,
		})
	}

	c.Locals("user_id", userID)
	c.Locals("user_role", userRole)
	return c.Next()
}
