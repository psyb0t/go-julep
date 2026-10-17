# go-julep

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-julep.svg)](https://pkg.go.dev/github.com/psyb0t/go-julep)
[![CI](https://github.com/psyb0t/go-julep/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-julep/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-julep/badges/version.svg)](https://github.com/psyb0t/go-julep/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-julep/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-julep/badges/importers.svg)](https://github.com/psyb0t/go-julep/blob/badges/importers.md)

Unofficial Go client for the [Julep API](https://docs.julep.ai), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Julep publishes. I am not affiliated with Julep.

The package is `julep`, in the repository root. Everything in `models.gen.go` and `client.gen.go` is generated. Don't edit them by hand, regenerate them.

## Install

```sh
go get github.com/psyb0t/go-julep
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200`, `JSON201` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-julep"
)

func main() {
	auth := julep.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("JULEP_API_KEY"))

		return nil
	})

	client, err := julep.NewClientWithResponses("https://api.julep.ai/api", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.ListAgentsAgentsGetWithResponse(context.Background(), &julep.ListAgentsAgentsGetParams{})
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

The full API, with the spec's descriptions as doc comments, is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-julep) and in `go doc github.com/psyb0t/go-julep`.

Method names come from the spec's operation IDs, which repeat the path and the HTTP method, so they are long: `ListAgentsAgentsGet` lists agents, `CreateTaskExecutionTasksTaskIdExecutionsPost` starts a task execution, `ChatSessionsSessionIdChatPost` sends a chat message in a session.

## Auth

The spec declares an API key in the `Authorization` header. Julep's official SDKs send it as `Authorization: Bearer <key>`, and the example above does the same. Keys come from the Julep dashboard.

| Environment | Base URL |
|---|---|
| Production | `https://api.julep.ai/api` |
| Dev | `https://dev.julep.ai/api` |

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` reads `openapi_spec_url` from the `.stats.yml` of [julep-ai/python-sdk](https://github.com/julep-ai/python-sdk) and downloads that URL to `api/openapi.yaml`. Stainless hosts the spec under a URL that contains a content hash, so the URL changes with every spec update.
- oapi-codegen runs twice. `config.yaml` writes the models to `models.gen.go` and `config.client.yaml` writes the client to `client.gen.go`, because one file would be over 40,000 lines. Both configs map `number` to `float64`, because the spec gives temperatures and other numbers no format and oapi-codegen would otherwise use `float32`.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-julep) teaches a coding agent how to use go-julep when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-julep@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-julep@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-julep
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Julep's published API specification. The specification itself is not redistributed here.
