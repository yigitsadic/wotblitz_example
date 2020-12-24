package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yigitsadic/wotblitz_example/shared"
	"log"

	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"github.com/yigitsadic/wotblitz_example/graph/model"
)

func (r *queryResolver) Tanks(ctx context.Context) ([]*model.Tank, error) {
	var tanks []*model.Tank

	foundTanks, err := r.Repository.FetchAllTanks()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {

		tanks = append(tanks, &model.Tank{
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: shared.MapTankClass(tank.TankClass),
			Country:   shared.MapTankCountry(tank.Country),
		})
	}

	return tanks, nil
}

func (r *queryResolver) TechTree(ctx context.Context, country model.Country) ([]*model.Tank, error) {
	var tanks []*model.Tank

	foundTanks, err := r.Repository.FetchTanksByCountry(country)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {

		tanks = append(tanks, &model.Tank{
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: shared.MapTankClass(tank.TankClass),
			Country:   shared.MapTankCountry(tank.Country),
		})
	}

	return tanks, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
