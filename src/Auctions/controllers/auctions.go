package controllers

import (
	"errors"
	"log/slog"
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

func CreateAuction(ctx *gin.Context) {
	/* TODO:
	Implement authentication on this route.
	*/

	var data dtos.CreateAuctionDTO

	if err := ctx.BindJSON(&data); err != nil {
		slog.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
		return
	}

	endDate, err := data.EndDate()
	if err != nil {
		// TODO: Don't show the raw error when DEBUG=0
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid End Date", "details": err.Error()})
		return
	}

	// TODO: Add current user as seller.
	seller := "test"

	auction := models.Auction{
		Item: models.Item{
			Make:     data.Make,
			Model:    data.Model,
			Color:    data.Color,
			Year:     data.Year,
			Mileage:  data.Mileage,
			ImageUrl: data.ImageUrl,
		},
		Seller:     seller,
		Status:     models.Live,
		AuctionEnd: endDate,
	}

	result := inits.DB.Create(&auction)
	if result.Error != nil {
		// TODO: Dig into the error types and do some better error handling.
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable To Create Auction", "details": result.Error.Error()})
		return
	}
	slog.Debug("Auction created.", "user", seller, "auctionId", auction.ID.String())
	auctionDTO := dtos.ToAuctionDTO(&auction)
	ctx.JSON(http.StatusCreated, gin.H{"data": auctionDTO})
}

func UpdateAuction(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid UUID"})
		return
	}

	var auction models.Auction
	result := inits.DB.Preload("Item").First(&auction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": err.Error()})
		}
	} else {
		// TODO: Check seller == user
		user := "test"

		var data dtos.UpdateAuctionDTO
		if err := ctx.BindJSON(&data); err != nil {
			slog.Warn(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
			return
		}

		if data.Make != "" {
			auction.Item.Make = data.Make
		}
		if data.Model != "" {
			auction.Item.Color = data.Color
		}
		if data.Mileage != 0 {
			auction.Item.Mileage = data.Mileage
		}
		if data.Color != "" {
			auction.Item.Color = data.Color
		}
		if data.Year != 0 {
			auction.Item.Year = data.Year
		}

		result := inits.DB.Save(&auction.Item)
		if result.Error != nil {
			// TODO: Dig into the error types and do some better error handling.
			slog.Error(result.Error.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable To Create Auction", "details": result.Error.Error()})
			return
		}
		slog.Debug("Auction updated.", "user", user, "auctionId", auction.ID.String())
		auctionDTO := dtos.ToAuctionDTO(&auction)
		ctx.JSON(http.StatusOK, gin.H{"data": auctionDTO})
	}
}

func DeleteAuction(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid UUID"})
		return
	}

	var auction models.Auction
	result := inits.DB.Preload("Item").First(&auction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": err.Error()})
		}
	} else {
		// TODO: Check seller == user
		user := "test"

		result := inits.DB.Delete(&auction)
		if result.Error != nil {
			// TODO: Dig into the error types and do some better error handling.
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable To Create Auction", "details": result.Error.Error()})
			return
		}
		slog.Debug("Auction deleted.", "user", user, "auctionId", rawId)
		ctx.Writer.WriteHeader(http.StatusNoContent)
	}
}
