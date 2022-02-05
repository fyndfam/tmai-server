package server

import (
	"github.com/fyndfam/tmai-server/src/controller"
	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewApp(env *env.Env) *fiber.App {
	app := fiber.New()

	controller.MountHealthCheckRoutes(env, app)
	controller.MountUserRoutes(env, app)
	controller.MountPostRoutes(env, app)

	app.Get("/", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(env), restricted)

	return app
}

func restricted(context *fiber.Ctx) error {
	return context.JSON(map[string]string{"status": "ok"})
}
