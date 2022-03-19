package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/service"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtMiddleware() fiber.Handler {
	var config jwtware.Config
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "production" || appEnv == "staging" {
		refreshRateLimit := time.Duration(5) * time.Minute
		refreshInterval := time.Minute

		config = jwtware.Config{
			SigningMethod:       "RS256",
			KeySetURL:           os.Getenv("JWKS_URI"),
			KeyRefreshInterval:  &refreshInterval,
			KeyRefreshRateLimit: &refreshRateLimit,
		}
	} else {
		config = jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    []byte("BSDGR3VVE3EHMTVEYRMTKSUB"),
		}
	}

	return jwtware.New(config)
}

// call this middleware after JWT middleware, this middleware will check for claims
// create user by email if it doesn't exists, or get the user from database if user exists
func GetPostJwtMiddleware(env *env.Env) fiber.Handler {
	var issuer, audience string
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "production" || appEnv == "staging" {
		issuer = os.Getenv("JWT_ISSUER")
		audience = os.Getenv("JWT_AUDIENCE")
	} else {
		issuer = "tmaiserver"
		audience = "tmaiserver"
	}

	return func(context *fiber.Ctx) error {
		token := context.Locals("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		if ok := claims.VerifyAudience(audience, true); !ok {
			context.Status(401).Send([]byte("Invalid token aud"))
			return nil
		}

		if ok := claims.VerifyIssuer(issuer, true); !ok {
			context.Status(401).Send([]byte("Invalid token iss"))
			return nil
		}

		externalUserId, ok := claims["sub"].(string)
		if !ok || len(externalUserId) == 0 {
			context.Status(401).Send([]byte("Invalid token sub"))
			return nil
		}

		user, err := service.GetUserByExternalUserId(env, externalUserId)
		if err != nil {
			log.Print(err)
		}

		if user == nil {
			var emailAddress string

			email, err := getUserEmailFromAuth0(token.Raw)
			if err != nil {
				log.Print(err)
				context.Status(502).Send([]byte("Internal server error"))
				return nil
			}

			emailAddress = email

			createdUser, err := service.CreateUser(env, emailAddress, externalUserId)

			if err != nil {
				log.Print(err)
				return err
			}

			user = createdUser
		}

		context.Locals("user", *user)
		context.Next()
		return nil
	}
}

type Auth0UserInfo struct {
	Email string `json:"email"`
}

func getUserEmailFromAuth0(token string) (string, error) {
	userInfoURL := os.Getenv("AUTH0_USER_INFO_URL")

	req, err := http.NewRequest("GET", userInfoURL, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", err
	}

	var auth0UserInfo Auth0UserInfo
	err = json.Unmarshal(body, &auth0UserInfo)
	if err != nil {
		log.Panicln("Error when parsing JSON:", err)
		return "", err
	}

	return auth0UserInfo.Email, nil
}
