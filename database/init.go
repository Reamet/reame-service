package database

import (
	"fmt"
	"os"
	"reame-service/database/model"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host       string
	Port       int
	User       string
	Password   string
	SchemaName string
}

type DatabaseHandler struct {
	DB *gorm.DB
}

var Database *DatabaseHandler

func InitDbConfig() {
	Database = new(DatabaseHandler)

	Database.InitialPostgresql()
	Database.InitialMigration()
}

func (dh *DatabaseHandler) InitialPostgresql() {

	dbPort, dbPortError := strconv.Atoi(os.Getenv("DB_PORT"))

	if dbPortError != nil {
		panic(dbPortError)
	}

	databaseConfig := DatabaseConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       dbPort,
		User:       os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		SchemaName: os.Getenv("DB_SCHEMA_NAME"),
	}

	var err error

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.SchemaName)

	dh.DB, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func (dh *DatabaseHandler) InitialMigration() {
	dh.DB.AutoMigrate(&model.Collection{})
}
