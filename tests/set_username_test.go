package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/stretchr/testify/assert"
)

func TestSetUsernameAPI(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("POST /users/username", func() {
		environment := SetupTests()
		app := server.NewApp(environment)

		g.After(func() {
			TearDownTests(environment)
		})

		g.BeforeEach(func() {
			ClearDB(environment)
		})

		g.It("should not be able to set username again", func() {
			GivenUserWithUsername(environment)

			req := httptest.NewRequest("POST", "/users/username", strings.NewReader(`{"username": "realname"}`))
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", Bearer)

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 403, res.StatusCode)

			body, _ := ioutil.ReadAll(res.Body)
			var result map[string]interface{}
			parseErr := json.Unmarshal(body, &result)
			if parseErr != nil {
				t.Error(parseErr)
			}

			assert.Equal(t, "username already exists", result["error"])
		})

		g.It("should be able to set username for the first time", func() {
			GivenUser(environment)

			req := httptest.NewRequest("POST", "/users/username", strings.NewReader(`{"username": "realname"}`))
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

			assert.Equal(t, "success", result["status"])
		})
	})
}
