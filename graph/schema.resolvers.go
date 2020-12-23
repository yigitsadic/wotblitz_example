package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"github.com/yigitsadic/wotblitz_example/graph/model"
)

func (r *queryResolver) Tanks(ctx context.Context) ([]*model.Tank, error) {
	tanks := []*model.Tank{
		{
			Name:      "Tiger II",
			Tier:      8,
			IsPremium: false,
			TankClass: model.TankClassHeavyTank,
			Country:   model.CountryGermany,
		},
		{
			Name:      "IS-3",
			Tier:      8,
			IsPremium: false,
			TankClass: model.TankClassHeavyTank,
			Country:   model.CountryUsrr,
		},
	}

	return tanks, nil
}

func (r *queryResolver) TechTree(ctx context.Context, country model.Country) ([]*model.Tank, error) {
	tanks := []*model.Tank{
		{
			Name:      "Tiger II",
			Tier:      8,
			IsPremium: false,
			TankClass: model.TankClassHeavyTank,
			Country:   model.CountryGermany,
		},
	}

	return tanks, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
