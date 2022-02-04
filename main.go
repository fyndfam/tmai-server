package main

import (
	"fmt"
	"log"

	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	fmt.Println("hello")

	app := fiber.New()

	app.Get("/", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(), restricted)

	app.Listen(":8088")
}

func restricted(context *fiber.Ctx) error {
	user := context.Locals("user").(*jwt.Token)
	log.Printf("claims are %+v", user)

	return context.JSON(map[string]string{"status": "ok"})
}
