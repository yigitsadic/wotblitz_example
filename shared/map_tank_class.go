package shared

import "github.com/yigitsadic/wotblitz_example/graph/model"

func MapTankClass(tankClassFromDb string) model.TankClass {
	var tankClass model.TankClass

	switch tankClassFromDb {
	case "light":
		tankClass = model.TankClassLightTank
	case "medium":
		tankClass = model.TankClassMediumTank
	case "heavy":
		tankClass = model.TankClassHeavyTank
	}

	return tankClass
}

func MapTankCountry(tankCountryFromDb string) model.Country {
	var country model.Country

	switch tankCountryFromDb {
	case "usa":
		country = model.CountryUsa
	case "germany":
		country = model.CountryGermany
	case "ussr":
		country = model.CountryUsrr
	case "uk":
		country = model.CountryUk
	case "japan":
		country = model.CountryJapan
	case "china":
		country = model.CountryChina
	case "france":
		country = model.CountryFrance
	case "european_nation":
		country = model.CountryEuropeanNation
	}

	return country
}
