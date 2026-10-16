package julep_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-julep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey  = "test-key"
	testAgentID = "6f1c2a4e-0b3d-4c5e-8f9a-1b2c3d4e5f60"
	testUserID  = "0a1b2c3d-4e5f-4061-8293-a4b5c6d7e8f9"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *julep.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := julep.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)

			return nil
		},
	)

	client, err := julep.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestListAgents(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/agents", r.URL.Path)
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"items":[{"id":"` + testAgentID +
				`","name":"helper","model":"model-x"}]}`,
		))
	})

	limit := 10

	res, err := client.ListAgentsAgentsGetWithResponse(
		context.Background(),
		&julep.ListAgentsAgentsGetParams{Limit: &limit},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Items, 1)

	agent := res.JSON200.Items[0]
	assert.Equal(t, testAgentID, agent.Id.String())
	assert.Equal(t, "helper", agent.Name)
	assert.Equal(t, "model-x", *agent.Model)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/users", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))

		var body julep.CreateUserRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "alice", *body.Name)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(
			`{"id":"` + testUserID + `","name":"alice"}`,
		))
	})

	name := "alice"

	res, err := client.CreateUserUsersPostWithResponse(
		context.Background(),
		julep.CreateUserRequest{Name: &name},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, res.StatusCode())
	require.NotNil(t, res.JSON201)

	assert.Equal(t, testUserID, res.JSON201.Id.String())
	assert.Equal(t, "alice", *res.JSON201.Name)
}
