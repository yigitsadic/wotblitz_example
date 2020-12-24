//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/yigitsadic/wotblitz_example/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository *database.Repository
}
