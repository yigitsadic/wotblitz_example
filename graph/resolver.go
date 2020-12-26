//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/yigitsadic/wotblitz_example/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
}
