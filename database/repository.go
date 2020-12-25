package database

import (
	"errors"
	"github.com/yigitsadic/wotblitz_example/graph/model"
	"github.com/yigitsadic/wotblitz_example/shared"
	"strings"
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

// Returns next tanks for given tank id
func (r *Repository) FetchNextTanks(tankId int) ([]*TankSchema, error) {
	var foundNextTanks []*TankSchema

	for _, item := range r.Schema.TankFollowings {
		if item.FromTankId == tankId {

			for _, item2 := range r.Schema.Tanks {
				if item2.Id == item.ToTankId {
					foundNextTanks = append(foundNextTanks, item2)
				}
			}

		}
	}

	return foundNextTanks, nil
}

// Returns previous tanks for given tank id
func (r *Repository) FetchPreviousTanks(tankId int) ([]*TankSchema, error) {
	var foundNextTanks []*TankSchema

	for _, item := range r.Schema.TankFollowings {
		if item.ToTankId == tankId {

			for _, item2 := range r.Schema.Tanks {
				if item2.Id == item.FromTankId {
					foundNextTanks = append(foundNextTanks, item2)
				}
			}

		}
	}

	return foundNextTanks, nil
}

// Fetches modules for given tank id.
func (r *Repository) FetchModulesForTank(tankId int) ([]*ModuleSchema, error) {
	var foundModules []*ModuleSchema

	for _, m := range r.Schema.TankModules {
		if m.TankId == tankId {
			for _, mm := range r.Schema.Modules {
				if m.ModuleId == mm.Id {
					foundModules = append(foundModules, mm)
				}
			}
		}
	}

	return foundModules, nil
}

// Fetches stock modules for given tank id.
func (r *Repository) FetchStockModulesForTank(tankId int) ([]*ModuleSchema, error) {
	var foundModules []*ModuleSchema

	for _, m := range r.Schema.TankModules {
		if m.TankId == tankId {
			for _, mm := range r.Schema.Modules {
				if m.ModuleId == mm.Id && m.IsStock {
					foundModules = append(foundModules, mm)
				}
			}
		}
	}

	return foundModules, nil
}

// Fetches upgrade modules for given tank id.
func (r *Repository) FetchUpgradeModulesForTank(tankId int) ([]*ModuleSchema, error) {
	var foundModules []*ModuleSchema

	for _, m := range r.Schema.TankModules {
		if m.TankId == tankId {
			for _, mm := range r.Schema.Modules {
				if m.ModuleId == mm.Id && !m.IsStock {
					foundModules = append(foundModules, mm)
				}
			}
		}
	}

	return foundModules, nil
}

// Searches given term in tanks.
func (r *Repository) SearchInTanks(term string) ([]*TankSchema, error) {
	var foundTanks []*TankSchema

	for _, tank := range r.Schema.Tanks {
		if strings.Contains(strings.ToLower(tank.Name), strings.ToLower(term)) {
			foundTanks = append(foundTanks, tank)
		}
	}

	return foundTanks, nil
}

// Searches given term in modules.
func (r *Repository) SearchInModules(term string) ([]*ModuleSchema, error) {
	var foundModules []*ModuleSchema

	for _, module := range r.Schema.Modules {
		if strings.Contains(strings.ToLower(module.Name), strings.ToLower(term)) {
			foundModules = append(foundModules, module)
		}
	}

	return foundModules, nil
}
