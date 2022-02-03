package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
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
