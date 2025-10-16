package dtos

import (
	"time"

	"github.com/google/uuid"

	"github.com/seancallaway/CarBay/auction/models"
)

type AuctionDTO struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	AuctionEnd     time.Time `json:"auctionEnd"`
	Seller         string    `json:"seller"`
	Winner         *string   `json:"winner"`
	Make           string    `json:"make"`
	Model          string    `json:"model"`
	Year           int       `json:"year"`
	Color          string    `json:"color"`
	Mileage        int       `json:"mileage"`
	ImageUrl       string    `json:"imageUrl"`
	Status         string    `json:"status"`
	ReservePrice   int       `json:"reservePrice"`
	SoldAmount     int       `json:"soldAmount"`
	CurrentHighBid int       `json:"currentHighBid"`
}

type CreateAuctionDTO struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	Mileage      int    `json:"mileage,string"`
	Year         int    `json:"year,string"`
	ReservePrice int    `json:"reservePrice"`
	ImageUrl     string `json:"imageUrl"`
	AuctionEnd   string `json:"auctionEnd"`
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

func (dto CreateAuctionDTO) EndDate() (time.Time, error) {
	return time.Parse(time.RFC3339, dto.AuctionEnd)
}
