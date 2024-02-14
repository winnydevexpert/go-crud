package database

import (
	"fmt"
	"go-crud-learn/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(runEnv string) (func(), error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.AppConfig.DatabaseHost,
		config.AppConfig.DatabasePort,
		config.AppConfig.DatabaseUsername,
		config.AppConfig.DatabasePassword,
		config.AppConfig.DatabaseName,
	)

	var errGormOpen error
	DB, errGormOpen = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if errGormOpen != nil {
		return func() {}, errGormOpen
	}

	closeConnect := func() {
		db, err := DB.DB()
		if err == nil {
			db.Close()
		}
	}

	return closeConnect, nil
}
