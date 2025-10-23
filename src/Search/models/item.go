package models

import (
	"time"
)

type Item struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
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
