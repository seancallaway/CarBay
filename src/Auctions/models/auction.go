package models

import (
	"time"

	"github.com/google/uuid"
)

type AuctionStatus int

const (
	Live AuctionStatus = iota
	Finished
	ReserveNotMet
)

type Auction struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	ReservePrice   int       `gorm:"default:0"`
	Seller         string    `gorm:"not null"`
	Winner         string
	SoldAmount     int
	CurrentHighBid int
	CreatedAt      time.Time     `gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `gorm:"autoUpdateTime"`
	AuctionEnd     time.Time     `gorm:"not null"`
	Status         AuctionStatus `gorm:"default:0"`
	ItemID         uuid.UUID     `gorm:"not null"`
	Item           Item          `gorm:"not null"`
}
