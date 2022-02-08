package server

import (
	"github.com/fyndfam/tmai-server/src/controller"
	"github.com/fyndfam/tmai-server/src/env"
	"github.com/gofiber/fiber/v2"
)

func NewApp(env *env.Env) *fiber.App {
	app := fiber.New()

	controller.MountHealthCheckRoutes(env, app)
	controller.MountUserRoutes(env, app)
	controller.MountPostRoutes(env, app)

	return app
}
