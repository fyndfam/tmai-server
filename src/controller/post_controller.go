package controller

import (
	"log"
	"strconv"
	"strings"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/service"
	"github.com/gofiber/fiber/v2"
)

type InsertPostInput struct {
	Content string `json:"content"`
}

func MountPostRoutes(env *env.Env, app *fiber.App) {
	app.Post("/posts", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(env), createInsertPostEndpoint(env))
	app.Get("/posts", createGetLatestPostsEndpoint(env))
	app.Get("/posts/:postId", createGetPostByIdEndpoint(env))
}

func createInsertPostEndpoint(env *env.Env) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(model.UserModel)
		if !ok {
			ctx.Status(401)
			return nil
		}
		if user.Username == nil {
			ctx.Status(403).JSON(map[string]string{"message": "Please set your username before trying to create a post"})
			return nil
		}

		var input InsertPostInput

		if err := ctx.BodyParser(&input); err != nil {
			log.Println("Error parsing input", err)
			ctx.Status(400)
			return nil
		}

		content := strings.TrimSpace(input.Content)
		if len(content) == 0 {
			log.Println("Can not create post with empty content")
			ctx.Status(400).JSON(map[string]string{"message": "Post content must not be empty"})
			return nil
		}

		createdPost, err := service.CreatePost(env, &user, content)
		if err != nil {
			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(createdPost)
		return nil
	}
}

func createGetLatestPostsEndpoint(env *env.Env) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		offset, convErr := strconv.Atoi(ctx.Query("offset", "0"))
		if convErr != nil {
			ctx.Status(400)
			return nil
		}

		latestPosts, err := service.GetLatestPosts(env, 30, int64(offset))
		if err != nil {
			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(latestPosts)
		return nil
	}
}

func createGetPostByIdEndpoint(env *env.Env) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Params("postId")

		post, _ := service.GetPostByID(env, postID)
		if post == nil {
			ctx.Status(404)
			return nil
		}

		service.IncrementPostView(env, postID)

		ctx.Status(200).JSON(post)
		return nil
	}
}
