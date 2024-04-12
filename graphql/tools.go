//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

//go run github.com/99designs/gqlgen init
// go run github.com/99designs/gqlgen generate
// https://gqlgen.com/
