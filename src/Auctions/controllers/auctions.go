package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seancallaway/CarBay/auction/dtos"
	"github.com/seancallaway/CarBay/auction/inits"
	"github.com/seancallaway/CarBay/auction/models"
)

func GetAuctions(ctx *gin.Context) {
	var auctions []models.Auction
	// inits.DB.Joins("Item").Preload("Item").Order("items.make ASC").Find(&auctions)
	inits.DB.Joins("Item").
		Preload("Item").
		Order(`"Item"."make" ASC`).
		Find(&auctions)

	auctionDTOs := dtos.ToAuctionDTOs(auctions)

	ctx.JSON(http.StatusOK, gin.H{"data": auctionDTOs})
}
