package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seancallaway/CarBay/auction/controllers"
	"github.com/seancallaway/CarBay/auction/inits"
	"github.com/seancallaway/CarBay/auction/internal/utils"
)

func main() {
	if utils.Getenv("DEBUG", "0") == "0" {
		slog.Info("Running in production mode.")
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	inits.ConnectDatabase()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/api/auctions", controllers.GetAuctions)
	r.GET("/api/auctions/:id", controllers.GetAuctionById)

	r.Run()
}
