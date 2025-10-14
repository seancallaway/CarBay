package models

import (
	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID `jsoon:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Make     string    `gorm:"not null"`
	Model    string    `gorm:"not null"`
	Year     int       `gorm:"not null"`
	Color    string    `gorm:"not null"`
	Mileage  int       `gorm:"not null"`
	ImageUrl string    `gorm:"not null"`
}
