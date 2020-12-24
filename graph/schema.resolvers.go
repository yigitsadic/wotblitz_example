package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"github.com/yigitsadic/wotblitz_example/graph/model"
	"github.com/yigitsadic/wotblitz_example/shared"
)

func (r *queryResolver) FilterTanks(ctx context.Context, country *model.Country, tier *int, tankClass *model.TankClass, isPremium *bool) ([]*model.Tank, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResult, error) {
	var searchResult []model.SearchResult

	searchResult = append(searchResult, model.Module{
		Type: model.ModuleTypeSuspension,
		Name: "Tracks",
	})

	searchResult = append(searchResult, model.Tank{
		ID:        3,
		Name:      "Tiger II",
		Tier:      8,
		IsPremium: false,
		TankClass: model.TankClassHeavyTank,
		Country:   model.CountryGermany,
	})

	return searchResult, nil
}

func (r *queryResolver) Tanks(ctx context.Context) ([]*model.Tank, error) {
	var tanks []*model.Tank

	foundTanks, err := r.Repository.FetchAllTanks()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {

		tanks = append(tanks, &model.Tank{
			ID:        tank.Id,
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: shared.MapTankClass(tank.TankClass),
			Country:   shared.MapTankCountry(tank.Country),
		})
	}

	return tanks, nil
}

func (r *queryResolver) Tank(ctx context.Context, name string) (*model.Tank, error) {
	tank, err := r.Repository.FetchTankByName(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.Tank{
		ID:        tank.Id,
		Name:      tank.Name,
		Tier:      tank.Tier,
		IsPremium: tank.IsPremium,
		TankClass: shared.MapTankClass(tank.TankClass),
		Country:   shared.MapTankCountry(tank.Country),
	}, nil
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
			ID:        tank.Id,
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: shared.MapTankClass(tank.TankClass),
			Country:   shared.MapTankCountry(tank.Country),
		})
	}

	return tanks, nil
}

func (r *tankResolver) Modules(ctx context.Context, obj *model.Tank, isStock *bool) ([]*model.Module, error) {
	// TODO: Implement is stock search.

	var modules []*model.Module
	foundModules, err := r.Repository.FetchModulesForTank(obj.ID)
	if err != nil {
		return nil, nil
	}

	for _, m := range foundModules {
		modules = append(modules, &model.Module{
			Type: shared.MapModuleType(m.Type),
			Name: m.Name,
		})
	}

	return modules, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tank returns generated.TankResolver implementation.
func (r *Resolver) Tank() generated.TankResolver { return &tankResolver{r} }

type queryResolver struct{ *Resolver }
type tankResolver struct{ *Resolver }
