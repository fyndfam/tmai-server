package controller

import (
	"log"
	"strings"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/service"
	"github.com/gofiber/fiber/v2"
)

type UpdateUsernameInput struct {
	Username string `json:"username"`
}

func MountUserRoutes(env *env.Env, app *fiber.App) {
	app.Get("/users", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(env), createGetUserEndpoint(env))
	app.Post("/users/username", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(env), createUpdateUsernameEndpoint(env))
}

func createGetUserEndpoint(env *env.Env) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(model.UserModel)
		if !ok {
			ctx.Status(401)
			return nil
		}

		ctx.Status(200).JSON(user)
		return nil
	}
}

func createUpdateUsernameEndpoint(env *env.Env) fiber.Handler {
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

		if err := service.UpdateUsername(env, username, user.Email); err != nil {
			if err.Error() == "USERNAME_EXISTS" {
				ctx.Status(403).JSON(map[string]string{"error": "username already exists"})
				return nil
			}

			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(map[string]string{"status": "success"})
		return nil
	}
}
