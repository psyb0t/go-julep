# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `julep`, a REST client generated with oapi-codegen from Julep's OpenAPI document (Julep Agents API, spec version 0.4.0, OpenAPI 3.1.0, hosted by Stainless for the official SDKs). It covers all 37 paths: agents, docs and search, embeddings, executions and transitions, files, health, jobs, projects, secrets, sessions and chat, tasks, Temporal payload decoding, tools and users.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- The generated code is split into `models.gen.go` and `client.gen.go`.
- `make generate` reads the current spec URL from the `.stats.yml` of Julep's Python SDK, downloads the spec and regenerates the client. The spec is not committed.
- Added the `go-julep` agent skill. Install it with `claude plugin install go-julep@psyb0t` (Claude Code), `codex plugin add go-julep@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-julep` (ClawHub).
