# Go code

Refines `~/.claude/rule-details/working-with-generated-code.md` and `~/.claude/rule-details/coding-style.md`.

- **GO1** `models.gen.go` and `client.gen.go` are generated. Never edit them by hand. Change `config.yaml`, `config.client.yaml`, a script in `scripts/`, or `gen.go`, then run `make generate`.
- **GO2** `config.yaml` (models) and `config.client.yaml` (client) keep the same `output-options` and `compatibility` settings, or the two files disagree on type names and the package stops compiling.
- **GO3** A spec problem that breaks oapi-codegen is fixed by an oapi-codegen config option or a patch script in `scripts/` that `gen.go` runs between the fetch and the generate steps. Never by editing the downloaded spec or the output.
- **GO4** The whole spec is generated. No `include-tags` or other trimming.
- **GO5** `gen.go` holds only the package doc comment, the package clause and `//go:generate` lines. No other hand-written code goes in the package.
- **GO6** `make lint` passes with zero issues on the generated and hand-written code. Lines in hand-written files are at most 80 columns, the `.golangci.yml` limit.
- **GO7** Scripts in `scripts/` keep the template shape: `set -euo pipefail`, the `on_error` trap, JSON `log` lines via jq, and they pass shellcheck and shfmt.
