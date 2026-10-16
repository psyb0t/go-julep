// Package julep is a Go client for the Julep API, generated with oapi-codegen
// from the OpenAPI document Julep publishes.
package julep

//go:generate bash scripts/fetch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.yaml
//go:generate go tool oapi-codegen --config=config.client.yaml api/openapi.yaml
