# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `julep`, a REST client generated with oapi-codegen from Julep's OpenAPI document (Julep Agents API, spec version 0.4.0, OpenAPI 3.1.0, hosted by Stainless for the official SDKs). It covers all 37 paths: agents, docs and search, embeddings, executions and transitions, files, health, jobs, projects, secrets, sessions and chat, tasks, Temporal payload decoding, tools and users.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- The generated code is split into `models.gen.go` and `client.gen.go`.
- `make generate` reads the current spec URL from the `.stats.yml` of Julep's Python SDK, downloads the spec and regenerates the client. The spec is not committed.
- Added the `go-julep` agent skill. Install it with `claude plugin install go-julep@psyb0t` (Claude Code), `codex plugin add go-julep@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-julep` (ClawHub).
