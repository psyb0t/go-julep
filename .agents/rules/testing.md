# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `julep_test`, and use only the exported API.
- **TST2** REST tests run against `httptest.NewServer` and assert the HTTP method, the path, the `Authorization: Bearer` header, and the decoded response fields. No test calls the live API.
- **TST3** Tests call `t.Parallel()` and use `require` for preconditions and `assert` for values.
- **TST4** `example_test.go` holds the README usage example. When the README example changes, the example changes in the same commit, so `make test` keeps proving it compiles.
- **TST5** Test data uses obvious placeholders, never real keys, IDs or customer data.
