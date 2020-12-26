package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	tank2 "github.com/yigitsadic/wotblitz_example/ent/tank"
	"log"

	"github.com/yigitsadic/wotblitz_example/database"
	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"github.com/yigitsadic/wotblitz_example/graph/model"
	"github.com/yigitsadic/wotblitz_example/shared"
)

func (r *queryResolver) FilterTanks(ctx context.Context, country *model.Country, tier *int, tankClass *model.TankClass, isPremium *bool) ([]*model.Tank, error) {
	if country != nil {
		log.Println(*country)
	}

	if tier != nil {
		log.Println(*tier)
	}

	if tankClass != nil {
		log.Println(*tankClass)
	}

	if isPremium != nil {
		log.Println(*isPremium)
	}

	return nil, nil
}

func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResult, error) {
	var searchResult []model.SearchResult

	foundTanks, err := r.Repository.SearchInTanks(term)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {
		searchResult = append(searchResult, model.Tank{
			ID:        tank.Id,
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: shared.MapTankClass(tank.TankClass),
			Country:   shared.MapTankCountry(tank.Country),
		})
	}

	foundModules, err := r.Repository.SearchInModules(term)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, module := range foundModules {
		searchResult = append(searchResult, model.Module{
			Type: shared.MapModuleType(module.Type),
			Name: module.Name,
		})
	}

	return searchResult, nil
}

func (r *queryResolver) Tanks(ctx context.Context) ([]*model.Tank, error) {
	var tanks []*model.Tank

	foundTanks, err := r.Client.Tank.Query().All(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {
		tanks = append(tanks, &model.Tank{
			ID:        tank.ID,
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: model.TankClass(tank.TankClass),
			Country:   model.Country(tank.Country),
		})
	}

	return tanks, nil
}

func (r *queryResolver) Tank(ctx context.Context, name string) (*model.Tank, error) {
	tank, err := r.Client.Tank.Query().Where(tank2.Name(name)).First(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.Tank{
		ID:        tank.ID,
		Name:      tank.Name,
		Tier:      tank.Tier,
		IsPremium: tank.IsPremium,
		TankClass: model.TankClass(tank.TankClass),
		Country:   model.Country(tank.Country),
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

func (r *tankResolver) NextTanks(ctx context.Context, obj *model.Tank) ([]*model.Tank, error) {
	var tanks []*model.Tank
	foundTanks, err := r.Repository.FetchNextTanks(obj.ID)
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

func (r *tankResolver) PreviousTanks(ctx context.Context, obj *model.Tank) ([]*model.Tank, error) {
	var tanks []*model.Tank
	foundTanks, err := r.Repository.FetchPreviousTanks(obj.ID)
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

func (r *tankResolver) Modules(ctx context.Context, obj *model.Tank, status *model.ModuleFilter) ([]*model.Module, error) {
	var modules []*model.Module
	var err error
	var foundModules []*database.ModuleSchema

	if status == nil {
		foundModules, err = r.Repository.FetchModulesForTank(obj.ID)
	} else if *status == model.ModuleFilterStock {
		foundModules, err = r.Repository.FetchStockModulesForTank(obj.ID)
	} else if *status == model.ModuleFilterUpgrade {
		foundModules, err = r.Repository.FetchUpgradeModulesForTank(obj.ID)
	} else {
		foundModules, err = r.Repository.FetchModulesForTank(obj.ID)
	}

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
