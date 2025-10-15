package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seancallaway/CarBay/auction/controllers"
	"github.com/seancallaway/CarBay/auction/inits"
)

func main() {
	r := gin.Default()

	inits.ConnectDatabase()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/api/auctions", controllers.GetAuctions)

	r.Run("127.0.0.1:7001")
}
