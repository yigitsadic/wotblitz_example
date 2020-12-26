package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yigitsadic/wotblitz_example/ent/module"
	"github.com/yigitsadic/wotblitz_example/ent/predicate"
	tank2 "github.com/yigitsadic/wotblitz_example/ent/tank"
	"log"

	"github.com/yigitsadic/wotblitz_example/graph/generated"
	"github.com/yigitsadic/wotblitz_example/graph/model"
)

func (r *queryResolver) FilterTanks(ctx context.Context, country *model.Country, tier *int, tankClass *model.TankClass, isPremium *bool) ([]*model.Tank, error) {
	var queryParams []predicate.Tank

	if country != nil {
		queryParams = append(queryParams, tank2.Country(country.String()))
	}

	if tier != nil {
		queryParams = append(queryParams, tank2.Tier(*tier))
	}

	if tankClass != nil {
		queryParams = append(queryParams, tank2.TankClass(tankClass.String()))
	}

	if isPremium != nil {
		queryParams = append(queryParams, tank2.IsPremium(*isPremium))
	}

	var tanks []*model.Tank
	foundTanks, err := r.Client.Tank.Query().Where(queryParams...).All(ctx)
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

func (r *queryResolver) Search(ctx context.Context, term string) ([]model.SearchResult, error) {
	var searchResult []model.SearchResult

	foundTanks, err := r.Client.Tank.Query().Where(tank2.NameContains(term)).All(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, tank := range foundTanks {
		searchResult = append(searchResult, model.Tank{
			ID:        tank.ID,
			Name:      tank.Name,
			Tier:      tank.Tier,
			IsPremium: tank.IsPremium,
			TankClass: model.TankClass(tank.TankClass),
			Country:   model.Country(tank.Country),
		})
	}

	foundModules, err := r.Client.Module.Query().Where(module.NameContains(term)).All(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, foundModule := range foundModules {
		searchResult = append(searchResult, model.Module{
			Type: model.ModuleType(foundModule.ModuleType),
			Name: foundModule.Name,
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

	foundTanks, err := r.Client.Tank.Query().Where(tank2.Country(country.String())).All(ctx)

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

func (r *tankResolver) NextTanks(ctx context.Context, obj *model.Tank) ([]*model.Tank, error) {
	var tanks []*model.Tank
	foundTanks, err := r.Client.Tank.Query().Where(tank2.ID(obj.ID)).QueryNextTanks().All(ctx)

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

func (r *tankResolver) PreviousTanks(ctx context.Context, obj *model.Tank) ([]*model.Tank, error) {
	var tanks []*model.Tank
	foundTanks, err := r.Client.Tank.Query().Where(tank2.ID(obj.ID)).QueryPreviousTanks().All(ctx)

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

func (r *tankResolver) Modules(ctx context.Context, obj *model.Tank, status *model.ModuleFilter) ([]*model.Module, error) {
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tank returns generated.TankResolver implementation.
func (r *Resolver) Tank() generated.TankResolver { return &tankResolver{r} }

type queryResolver struct{ *Resolver }
type tankResolver struct{ *Resolver }
