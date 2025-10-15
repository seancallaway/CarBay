package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/seancallaway/CarBay/auction/dtos"
	"github.com/seancallaway/CarBay/auction/inits"
	"github.com/seancallaway/CarBay/auction/models"
)

func GetAuctions(ctx *gin.Context) {
	var auctions []models.Auction
	inits.DB.Joins("Item").
		Preload("Item").
		Order(`"Item"."make" ASC`).
		Find(&auctions)

	auctionDTOs := dtos.ToAuctionDTOs(auctions)

	ctx.JSON(http.StatusOK, gin.H{"data": auctionDTOs})
}

func GetAuctionById(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var auction models.Auction
	result := inits.DB.Preload("Item").First(&auction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		}
	} else {
		auctionDTO := dtos.ToAuctionDTO(&auction)
		ctx.JSON(http.StatusOK, gin.H{"data": auctionDTO})
	}
}
