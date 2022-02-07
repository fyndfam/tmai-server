package tests

import (
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestSetUsernameAPI(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("POST /users/username", func() {
		g.Before(func() {
			SetupTests()
		})

		g.After(func() {
			TearDownTests()
		})

		g.BeforeEach(func() {
			ClearDB()
		})

		g.It("should not be able to set username again", func() {
			GivenUserWithUsername()

			apitest.New().
				HandlerFunc(FiberToHandler(server.NewApp(GetEnvironment()))).
				Post("/users/username").
				Header("Authorization", Bearer).
				Header("Content-type", "application/json").
				Body(`{"username": "realname"}`).
				Expect(t).
				Assert(jsonpath.Equal(`$.error`, "username already exists")).
				Status(http.StatusForbidden).
				End()
		})

		g.It("should be able to set username for the first time", func() {
			GivenUser()

			apitest.New().
				HandlerFunc(FiberToHandler(server.NewApp(GetEnvironment()))).
				Post("/users/username").
				Header("Authorization", Bearer).
				Header("Content-type", "application/json").
				Body(`{"username": "realname"}`).
				Expect(t).
				Assert(jsonpath.Equal(`$.status`, "success")).
				Status(http.StatusOK).
				End()
		})
	})
}
