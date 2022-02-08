package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/fyndfam/tmai-server/src/server"
	"github.com/stretchr/testify/assert"
)

func TestGetPostById(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GET /posts/:postId", func() {
		environment := SetupTests()
		app := server.NewApp(environment)

		g.After(func() {
			TearDownTests(environment)
		})

		g.BeforeEach(func() {
			ClearDB(environment)
		})

		g.It("should be able to get post by ID", func() {
			GivenUserWithUsername(environment)
			postID := GivenPost(environment, "this is a sample post")

			req := httptest.
				NewRequest("GET", fmt.Sprintf("/posts/%v", postID), nil)
			req.Header.Set("Content-type", "application/json")

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
			assert.Equal(t, "test", result.CreatedBy)
		})
	})
}
