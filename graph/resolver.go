//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/yigitsadic/wotblitz_example/database"
	"github.com/yigitsadic/wotblitz_example/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository *database.Repository
	Client     *ent.Client
}
