package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtMiddleware() fiber.Handler {
	var config jwtware.Config

	if os.Getenv("APP_ENV") == "production" {
		refreshRateLimit := time.Duration(5) * time.Minute
		refreshInterval := time.Minute

		config = jwtware.Config{
			SigningMethod:       "RS256",
			KeySetURL:           os.Getenv("JWKS_URI"),
			KeyRefreshInterval:  &refreshInterval,
			KeyRefreshRateLimit: &refreshRateLimit,
		}
	} else {
		log.Println("using a dev setup for jwt config")

		config = jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    []byte("BSDGR3VVE3EHMTVEYRMTKSUB"),
		}
	}

	return jwtware.New(config)
}

// call this middleware after JWT middleware, this middleware will check for claims
// create user by email if it doesn't exists, or get the user from database if user exists
func GetPostJwtMiddleware() fiber.Handler {
	var issuer, audience string

	if os.Getenv("APP_ENV") == "production" {
		issuer = os.Getenv("JWT_ISSUER")
		audience = os.Getenv("JWT_AUDIENCE")
	} else {
		issuer = "tmaiserver"
		audience = "tmaiserver"
	}

	return func(context *fiber.Ctx) error {
		user := context.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		if ok := claims.VerifyAudience(issuer, true); !ok {
			context.Status(401).Send([]byte("Invalid token"))
			return nil
		}

		if ok := claims.VerifyIssuer(audience, true); !ok {
			context.Status(401).Send([]byte("Invalid token"))
			return nil
		}

		// TODO: get or create user by email

		context.Next()
		return nil
	}
}
