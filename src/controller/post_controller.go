package controller

import (
	"log"
	"strconv"
	"strings"

	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/service"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertPostInput struct {
	Content string `json:"content"`
}

func MountPostRoutes(mongoClient *mongo.Client, app *fiber.App) {
	app.Post("/posts", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(mongoClient), createInsertPostEndpoint(mongoClient))
	app.Get("/posts", createGetLatestPostsEndpoint(mongoClient))
	app.Get("/posts/:postId", createGetPostByIdEndpoint(mongoClient))
}

func createInsertPostEndpoint(mongoClient *mongo.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, ok := ctx.Locals("user").(model.UserModel)
		if !ok {
			ctx.Status(401)
			return nil
		}
		if user.Username == nil {
			ctx.Status(403)
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
			ctx.Status(400)
			return nil
		}

		createdPost, err := service.CreatePost(mongoClient, *user.Username, content)
		if err != nil {
			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(createdPost)
		return nil
	}
}

func createGetLatestPostsEndpoint(mongoClient *mongo.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		offset, convErr := strconv.Atoi(ctx.Query("offset", "0"))
		if convErr != nil {
			ctx.Status(400)
			return nil
		}

		latestPosts, err := service.GetLatestPosts(mongoClient, int64(offset), 0)
		if err != nil {
			ctx.Status(502)
			return nil
		}

		ctx.Status(200).JSON(latestPosts)
		return nil
	}
}

func createGetPostByIdEndpoint(mongoClient *mongo.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		postID := ctx.Params("postId")

		post, _ := service.GetPostByID(mongoClient, postID)
		if post == nil {
			ctx.Status(404)
			return nil
		}

		ctx.Status(200).JSON(post)
		return nil
	}
}
