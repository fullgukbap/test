package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-user/dto"
	"simple-user/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	app := fiber.New()
	app.Post("/auth/login", Login)
	app.Post("/auth/valid", jwtValidMiddleware, Valid)

	log.Fatal(app.Listen(":9190"))
}

// [POST] /auth/login
func Login(c *fiber.Ctx) error {
	// 파싱
	body := new(dto.LoginUserRequest)
	c.BodyParser(body)

	// 데이터베이스 비교 대조
	if _, err := repository.Read(body.Username); err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	// 토큰 발급
	token, err := GenerateToken(body.Username)
	if err != nil {
		fmt.Println(err)
	}

	// 반환
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func GenerateToken(username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": username,
		"exp": time.Now().Add(time.Second * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte("haha"))
}

func ValidToken(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("haha"), nil
	})

	var username string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = claims["iss"].(string)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", fiber.NewError(fiber.StatusUnauthorized)
		}
	} else {
		fmt.Println(err)
	}

	return username, nil
}

// [POST] /auth/validate
func jwtValidMiddleware(c *fiber.Ctx) error {
	body := struct {
		Token string `json:"token"`
	}{}

	c.BodyParser(&body)

	username, err := ValidToken(body.Token)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("username", username)
	return c.Next()
}

func Valid(c *fiber.Ctx) error {
	u := c.Locals("username").(string)
	return c.SendString(u)
}
