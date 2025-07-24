package graph

import "github.com/stack-boilerplate/go-graphql-gqlgen/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
}
