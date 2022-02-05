package main

import (
	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/server"
)

func main() {
	env := env.InitializeEnvironment()

	app := server.NewApp(env)

	app.Listen(":8088")
}
