package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/seancallaway/CarBay/search/database"
	"github.com/seancallaway/CarBay/search/internal/utils"
	"github.com/seancallaway/CarBay/search/models"
)

func main() {
	if utils.Getenv("DEBUG", "0") == "0" {
		slog.Info("Running in production mode.")
		gin.SetMode(gin.ReleaseMode)
	}

	DB := database.ConnectDB().Database("SearchDB")
	models.CreateItemIndexes(DB)
	err := models.InitDb(DB)
	if err != nil {
		slog.Error("Error seeding DB: " + err.Error())
	}

	r := gin.Default()
	r.Use(database.MongoDBMiddleware(DB))

	r.Run()
}
