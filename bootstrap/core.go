package bootstrap

import (
	"web/controller/resource"
	"web/data"
	"web/dataprovider/database"
	"web/provider/authservice"
	"web/provider/vocabulary"
)

func Init() {
	vocabulary.Init()
	authservice.Init()

	loadCore()
	loadResources()
}

func loadCore() {
	database.Init(data.DBConnection())
}

func loadResources() {
	resource.Init()

