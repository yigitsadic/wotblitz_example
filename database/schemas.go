package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TankSchema struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Tier      int    `json:"tier"`
	IsPremium bool   `json:"isPremium"`
	TankClass string `json:"tankClass"`
	Country   string `json:"country"`
}

type ModuleSchema struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	IsStock bool   `json:"-"`
}

type ModuleFollowingSchema struct {
	Id                int `json:"id"`
	ModuleId          int `json:"moduleId"`
	FollowingModuleId int `json:"followingModuleId"`
}

type TankFollowingSchema struct {
	Id         int `json:"id"`
	FromTankId int `json:"fromTankId"`
	ToTankId   int `json:"toTankId"`
}

type TankModuleSchema struct {
	Id       int  `json:"id"`
	TankId   int  `json:"tankId"`
	ModuleId int  `json:"moduleId"`
	IsStock  bool `json:"isStock"`
}

type Schema struct {
	Tanks            []*TankSchema            `json:"tanks"`
	Modules          []*ModuleSchema          `json:"modules"`
	ModuleFollowings []*ModuleFollowingSchema `json:"moduleFollowings"`
	TankModules      []*TankModuleSchema      `json:"tankModules"`
	TankFollowings   []*TankFollowingSchema   `json:"tankFollowings"`
}

func ReadFromFile() []byte {
	db, err := ioutil.ReadFile("database/database.json")
	if err != nil {
		log.Println(err)
		log.Fatal("unable to read database file")
	}

	return db
}

func LoadFromFile() *Schema {
	var schema Schema

	err := json.Unmarshal(ReadFromFile(), &schema)
	if err != nil {
		log.Println(err)
		log.Fatal("Unable to unmarshal json to struct")
	}

	return &schema
}
