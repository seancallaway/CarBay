package main

import (
	"github.com/seancallaway/CarBay/auction/inits"
	"github.com/seancallaway/CarBay/auction/models"
)

func init() {
	inits.ConnectDatabase()
}

func main() {
	err := inits.DB.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}
	err = inits.DB.AutoMigrate(&models.Auction{})
	if err != nil {
		panic(err)
	}
}
