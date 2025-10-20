package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/seancallaway/CarBay/search/database"
	"github.com/seancallaway/CarBay/search/internal/utils"
)

func main() {
	if utils.Getenv("DEBUG", "0") == "0" {
		slog.Info("Running in production mode.")
		gin.SetMode(gin.ReleaseMode)
	}

	DB := database.ConnectDB()

	r := gin.Default()
	r.Use(database.MongoDBMiddleware(DB))

	r.Run()
}
