package controller

import (
	"log"
	"strings"

	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateUsernameInput struct {
	Username string `json:"username"`
}

func MountUserRoutes(mongoClient *mongo.Client, app *fiber.App) {
	app.Post("/users/username", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(mongoClient), createUpdateUsernameEndpoint(mongoClient))
}

func createUpdateUsernameEndpoint(mongoClient *mongo.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(model.UserModel)
		if !ok {
			ctx.Status(401)
			return nil
		}

		var input UpdateUsernameInput

		if err := ctx.BodyParser(&input); err != nil {
			log.Println("Error parsing input", err)
			ctx.Status(400)
			return nil
		}

		username := strings.TrimSpace(input.Username)
		if len(username) < 2 || len(username) > 19 {
			log.Println("Invalid length for username")
			ctx.Status(400)
			return nil
		}

		log.Println("user email address is", user.Email)
		if err := service.UpdateUsername(mongoClient, username, user.Email); err != nil {
			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(map[string]string{"status": "success"})
		return nil
	}
}
