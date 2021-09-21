package database

import (
	"fmt"
	"gobin/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	var databaseDialector gorm.Dialector
	DBDriver := strings.ToLower(os.Getenv("DB_DRIVER"))
	if DBDriver == "" {
		DBDriver = "postgres"
	}
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBDatabase := os.Getenv("DB_DATABASE")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	switch DBDriver {
	case "sqlite":
		databaseDialector = sqlite.Open(DBDatabase)
	case "mysql":
		if DBUser == "" {
			DBUser = "localhost"
		}
		if DBPort == "" {
			DBPort = "3306"
		}
		if DBUser == "" {
			DBUser = "root"
		}
		databaseDialector = mysql.Open(fmt.Sprintf("%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBDatabase))
	case "postgres":
		if DBHost == "" {
			DBHost = "localhost"
		}
		if DBPort == "" {
			DBPort = "5432"
		}
		if DBUser == "" {
			DBUser = "postgres"
		}
		databaseDialector = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DBHost, DBUser, DBPassword, DBDatabase, DBPort))
	default:
		panic("Database Driver '" + DBDriver + "' is unknown. Please use one from sqlite,mysql or postgres")
	}

	DB, err = gorm.Open(databaseDialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to Database")
	err = DB.AutoMigrate(&model.PasteModel{})
	if err != nil {
		panic("failed to migrate database")
	}

	log.Println("Database Migrated")
}
