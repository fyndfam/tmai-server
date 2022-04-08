package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("POST /posts", func() {
		environment := SetupTests()
		app := server.NewApp(environment)

		g.After(func() {
			TearDownTests(environment)
		})

		g.BeforeEach(func() {
			ClearDB(environment)
		})

		g.It("should be able to post", func() {
			GivenUserWithUsername(environment)

			req := httptest.NewRequest("POST", "/posts", strings.NewReader(`{"content": "this is a sample post"}`))
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", Bearer)

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 200, res.StatusCode)

			body, _ := ioutil.ReadAll(res.Body)
			var result model.PostModel
			parseErr := json.Unmarshal(body, &result)
			if parseErr != nil {
				t.Error(parseErr)
			}

			assert.Equal(t, "this is a sample post", result.Content)
			assert.Equal(t, "test", result.CreatedBy.Username)
			assert.Equal(t, "avatar/avatar_1.png", result.CreatedBy.Avatar)
		})

		g.It("should return 403 if username is not set", func() {
			GivenUser(environment)

			req := httptest.NewRequest("POST", "/posts", strings.NewReader(`{"content": "this is a sample post"}`))
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", Bearer)

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 403, res.StatusCode)
		})

		g.It("should not be able to create post if not authenticated", func() {
			GivenUserWithUsername(environment)
			req := httptest.NewRequest("POST", "/posts", strings.NewReader(`{"content": "this is a sample post"}`))
			req.Header.Set("Content-type", "application/json")

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 400, res.StatusCode)
		})
	})
}
