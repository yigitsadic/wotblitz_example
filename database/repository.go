package database

import (
	"errors"
	"github.com/yigitsadic/wotblitz_example/graph/model"
	"github.com/yigitsadic/wotblitz_example/shared"
)

type Repository struct {
	Schema *Schema
}

func NewRepository(schema *Schema) *Repository {
	return &Repository{Schema: schema}
}

// Returns all tanks without any filter.
func (r *Repository) FetchAllTanks() ([]*TankSchema, error) {
	return r.Schema.Tanks, nil
}

// Returns all tanks belong to country from parameter.
func (r *Repository) FetchTanksByCountry(countryName model.Country) ([]*TankSchema, error) {
	var tanks []*TankSchema
	for _, tank := range r.Schema.Tanks {
		if shared.MapTankCountry(tank.Country) == countryName {
			tanks = append(tanks, tank)
		}
	}

	return tanks, nil
}

// Returns tank with matched name.
func (r *Repository) FetchTankByName(name string) (*TankSchema, error) {
	for _, tank := range r.Schema.Tanks {
		if tank.Name == name {
			return tank, nil
		}
	}

	return nil, errors.New("tank not found")
}

func (r *Repository) FetchTanksByIds(ids []int) ([]*TankSchema, error) {
	return nil, nil
}

func (r *Repository) FetchModulesByIds(ids []int) ([]*ModuleSchema, error) {
	return nil, nil
}

func (r *Repository) FetchStockModules(tankId int) ([]*ModuleSchema, error) {
	return nil, nil
}

func (r *Repository) FetchNextTanks(tankId int) ([]*TankSchema, error) {
	return nil, nil
}

func (r *Repository) FetchModulesForTank(tankId int) ([]*ModuleSchema, error) {
	return nil, nil
}
