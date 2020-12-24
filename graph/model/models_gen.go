// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Module struct {
	Type       ModuleType `json:"type"`
	Name       string     `json:"name"`
	Followings []*Module  `json:"followings"`
}

type Tank struct {
	Name      string    `json:"name"`
	Tier      int       `json:"tier"`
	NextTanks []*Tank   `json:"nextTanks"`
	Modules   []*Module `json:"modules"`
	IsPremium bool      `json:"isPremium"`
	TankClass TankClass `json:"tankClass"`
	Country   Country   `json:"country"`
}

type Country string

const (
	CountryUsa            Country = "USA"
	CountryGermany        Country = "Germany"
	CountryUsrr           Country = "USRR"
	CountryUk             Country = "UK"
	CountryJapan          Country = "Japan"
	CountryChina          Country = "China"
	CountryFrance         Country = "France"
	CountryEuropeanNation Country = "EuropeanNation"
)

var AllCountry = []Country{
	CountryUsa,
	CountryGermany,
	CountryUsrr,
	CountryUk,
	CountryJapan,
	CountryChina,
	CountryFrance,
	CountryEuropeanNation,
}

func (e Country) IsValid() bool {
	switch e {
	case CountryUsa, CountryGermany, CountryUsrr, CountryUk, CountryJapan, CountryChina, CountryFrance, CountryEuropeanNation:
		return true
	}
	return false
}

func (e Country) String() string {
	return string(e)
}

func (e *Country) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Country(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Country", str)
	}
	return nil
}

func (e Country) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ModuleType string

const (
	ModuleTypeTurret     ModuleType = "Turret"
	ModuleTypeGun        ModuleType = "Gun"
	ModuleTypeEngine     ModuleType = "Engine"
	ModuleTypeSuspension ModuleType = "Suspension"
)

var AllModuleType = []ModuleType{
	ModuleTypeTurret,
	ModuleTypeGun,
	ModuleTypeEngine,
	ModuleTypeSuspension,
}

func (e ModuleType) IsValid() bool {
	switch e {
	case ModuleTypeTurret, ModuleTypeGun, ModuleTypeEngine, ModuleTypeSuspension:
		return true
	}
	return false
}

func (e ModuleType) String() string {
	return string(e)
}

func (e *ModuleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ModuleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ModuleType", str)
	}
	return nil
}

func (e ModuleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TankClass string

const (
	TankClassTankDestroyer TankClass = "TankDestroyer"
	TankClassLightTank     TankClass = "LightTank"
	TankClassMediumTank    TankClass = "MediumTank"
	TankClassHeavyTank     TankClass = "HeavyTank"
)

var AllTankClass = []TankClass{
	TankClassTankDestroyer,
	TankClassLightTank,
	TankClassMediumTank,
	TankClassHeavyTank,
}

func (e TankClass) IsValid() bool {
	switch e {
	case TankClassTankDestroyer, TankClassLightTank, TankClassMediumTank, TankClassHeavyTank:
		return true
	}
	return false
}

func (e TankClass) String() string {
	return string(e)
}

func (e *TankClass) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TankClass(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TankClass", str)
	}
	return nil
}

func (e TankClass) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
