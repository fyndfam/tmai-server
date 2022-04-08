package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestPosts(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GET /posts", func() {
		environment := SetupTests()
		app := server.NewApp(environment)

		g.After(func() {
			TearDownTests(environment)
		})

		g.BeforeEach(func() {
			ClearDB(environment)
		})

		g.It("should be able to get latest posts", func() {
			GivenUserWithUsername(environment)
			GivenPosts(environment)

			req := httptest.NewRequest("GET", "/posts", nil)
			req.Header.Set("Content-type", "application/json")

			res, err := app.Test(req)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, 200, res.StatusCode)

			body, _ := ioutil.ReadAll(res.Body)
			var result []model.PostModel
			if parseErr := json.Unmarshal(body, &result); parseErr != nil {
				t.Error(parseErr)
			}

			assert.Equal(t, 2, len(result))

			firstPost := result[0]
			assert.Equal(t, "oh, this is interesting, i want it!", firstPost.Content)

			secondPost := result[1]
			assert.Equal(t, "hey, this is the first post", secondPost.Content)
		})
	})
}
