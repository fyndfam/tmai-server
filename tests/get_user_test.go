package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/stretchr/testify/assert"
)

func TestGetUserAPI(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GET /users", func() {
		environment := SetupTests()
		app := server.NewApp(environment)

		g.After(func() {
			TearDownTests(environment)
		})

		g.BeforeEach(func() {
			ClearDB(environment)
		})

		g.It("should be able to get user", func() {
			req := httptest.NewRequest("GET", "/users", nil)
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", Bearer)

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 200, res.StatusCode)

			body, _ := ioutil.ReadAll(res.Body)
			var result map[string]interface{}
			parseErr := json.Unmarshal(body, &result)
			if parseErr != nil {
				t.Error(parseErr)
			}

			assert.Equal(t, "test@tmai.co", result["email"])
		})
	})
}
