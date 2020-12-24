package database

type Repository struct {
	Schema *Schema
}

func (r *Repository) FetchTankById(id int) (*TankSchema, error) {
	return nil, nil
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
