package dtos

import (
	"time"

	"github.com/google/uuid"

	"github.com/seancallaway/CarBay/auction/models"
)

type AuctionDTO struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	AuctionEnd     time.Time `json:"auction_end"`
	Seller         string    `json:"seller"`
	Winner         *string   `json:"winner"`
	Make           string    `json:"make"`
	Model          string    `json:"model"`
	Year           int       `json:"year"`
	Color          string    `json:"color"`
	Mileage        int       `json:"mileage"`
	ImageUrl       string    `json:"image_url"`
	Status         string    `json:"status"`
	ReservePrice   int       `json:"reserve_price"`
	SoldAmount     int       `json:"sold_amount"`
	CurrentHighBid int       `json:"current_high_bid"`
}

func ToAuctionDTO(auction *models.Auction) AuctionDTO {
	auctionDTO := AuctionDTO{
		ID:             auction.ID,
		CreatedAt:      auction.CreatedAt,
		UpdatedAt:      auction.UpdatedAt,
		AuctionEnd:     auction.AuctionEnd,
		Seller:         auction.Seller,
		Winner:         auction.Winner,
		ReservePrice:   auction.ReservePrice,
		SoldAmount:     auction.SoldAmount,
		CurrentHighBid: auction.CurrentHighBid,
	}

	switch auction.Status {
	case models.Finished:
		auctionDTO.Status = "Finished"
	case models.ReserveNotMet:
		auctionDTO.Status = "Reserve Not Met"
	default:
		auctionDTO.Status = "Live"
	}

	if auction.Item.ID != uuid.Nil {
		auctionDTO.Make = auction.Item.Make
		auctionDTO.Model = auction.Item.Model
		auctionDTO.Color = auction.Item.Color
		auctionDTO.Year = auction.Item.Year
		auctionDTO.Mileage = auction.Item.Mileage
		auctionDTO.ImageUrl = auction.Item.ImageUrl
	}

	return auctionDTO
}

func ToAuctionDTOs(auctions []models.Auction) []AuctionDTO {
	var auctionDTOs []AuctionDTO
	for _, auction := range auctions {
		auctionDTOs = append(auctionDTOs, ToAuctionDTO(&auction))
	}
	return auctionDTOs
}
