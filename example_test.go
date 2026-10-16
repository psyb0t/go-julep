package julep_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-julep"
)

// This is the README usage example. It talks to the live API, so it has no
// Output comment and only has to compile.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := julep.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("JULEP_API_KEY"),
			)

			return nil
		},
	)

	client, err := julep.NewClientWithResponses(
		"https://api.julep.ai/api",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.ListAgentsAgentsGetWithResponse(
		context.Background(),
		&julep.ListAgentsAgentsGetParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
