package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type databaseConfiguration struct {
	Username               string
	Password               string
	Database               string
	Address                string
	PrivateAddress         string
	Port                   string
	CloudSQLConnectionName string
}

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbName,
	)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func ConnectDB() error {
	var err error
	DB, err = InitDB()

	return err
}
