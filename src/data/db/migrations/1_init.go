package migrations

import (
	"github.com/alirzamehrzad/drivey/data/db"
	"github.com/alirzamehrzad/drivey/data/models"
)

func Up_1() {
	database := db.GetDb()

	tabels := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tabels = append(tabels, country)
	}

	if !database.Migrator().HasTable(city) {
		tabels = append(tabels, city)
	}

	database.Migrator().CreateTable(tabels...)
}

func Down_1() {

}
