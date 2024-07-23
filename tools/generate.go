package tools

// OAPI codegen
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen.types.yaml  ../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen.client.yaml ../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config oapi-codegen.server.yaml ../openapi.yaml

// SQLC
//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate
