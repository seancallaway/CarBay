package inits

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/seancallaway/CarBay/auction/internal/utils"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUser := utils.Getenv("SQL_USER", "auction")
	dbPass := os.Getenv("SQL_PASSWORD")
	dbHost := utils.Getenv("SQL_HOST", "db")
	dbPort := utils.Getenv("SQL_PORT", "5432")
	dbName := utils.Getenv("SQL_NAME", "auction")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database! " + err.Error())
	}

	DB = database
}
