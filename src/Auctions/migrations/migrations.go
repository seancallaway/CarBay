package main

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/seancallaway/CarBay/auction/inits"
	"github.com/seancallaway/CarBay/auction/models"
)

func seedData() {
	var count int64
	inits.DB.Model(&models.Auction{}).Count(&count)
	if count > 0 {
		slog.Info("Database already contains data. No need to seed.")
		return
	} else {
		slog.Info("Nothing in the DB. Seeding data.")
	}

	auctions := []*models.Auction{
		{
			ID:           uuid.MustParse("afbee524-5972-4075-8800-7d1f9d7b0a0c"),
			Status:       models.Live,
			ReservePrice: 20000,
			Seller:       "bob",
			AuctionEnd:   time.Now().Add(10 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Ford",
				Model:    "GT",
				Color:    "White",
				Mileage:  50000,
				Year:     2020,
				ImageUrl: "https://cdn.pixabay.com/photo/2016/05/06/16/32/car-1376190_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("c8c3ec17-01bf-49db-82aa-1ef80b833a9f"),
			Status:       models.Live,
			ReservePrice: 90000,
			Seller:       "alice",
			AuctionEnd:   time.Now().Add(60 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Bugatti",
				Model:    "Veyron",
				Color:    "Black",
				Mileage:  15035,
				Year:     2018,
				ImageUrl: "https://cdn.pixabay.com/photo/2012/05/29/00/43/car-49278_960_720.jpg",
			},
		},
		{
			ID:         uuid.MustParse("bbab4d5a-8565-48b1-9450-5ac2a5c4a654"),
			Status:     models.Live,
			Seller:     "bob",
			AuctionEnd: time.Now().Add(4 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Ford",
				Model:    "Mustang",
				Color:    "Black",
				Mileage:  65125,
				Year:     2023,
				ImageUrl: "https://cdn.pixabay.com/photo/2012/11/02/13/02/car-63930_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("155225c1-4448-4066-9886-6786536e05ea"),
			Status:       models.ReserveNotMet,
			ReservePrice: 50000,
			Seller:       "tom",
			AuctionEnd:   time.Now().Add(-4 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Mercedes",
				Model:    "SLK",
				Color:    "Silver",
				Mileage:  15001,
				Year:     2020,
				ImageUrl: "https://cdn.pixabay.com/photo/2016/04/17/22/10/mercedes-benz-1335674_960_720.png",
			},
		},
		{
			ID:           uuid.MustParse("466e4744-4dc5-4987-aae0-b621acfc5e39"),
			Status:       models.Live,
			ReservePrice: 20000,
			Seller:       "alice",
			AuctionEnd:   time.Now().Add(30 * 24 * time.Hour),
			Item: models.Item{
				Make:     "BMW",
				Model:    "X1",
				Color:    "White",
				Mileage:  90000,
				Year:     2017,
				ImageUrl: "https://cdn.pixabay.com/photo/2017/08/31/05/47/bmw-2699538_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("dc1e4071-d19d-459b-b848-b5c3cd3d151f"),
			Status:       models.Live,
			ReservePrice: 20000,
			Seller:       "bob",
			AuctionEnd:   time.Now().Add(45 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Ferrari",
				Model:    "Spider",
				Color:    "Red",
				Mileage:  50000,
				Year:     2015,
				ImageUrl: "https://cdn.pixabay.com/photo/2017/11/09/01/49/ferrari-458-spider-2932191_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("47111973-d176-4feb-848d-0ea22641c31a"),
			Status:       models.Live,
			ReservePrice: 150000,
			Seller:       "alice",
			AuctionEnd:   time.Now().Add(13 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Ferrari",
				Model:    "F-430",
				Color:    "Red",
				Mileage:  5000,
				Year:     2022,
				ImageUrl: "https://cdn.pixabay.com/photo/2017/11/08/14/39/ferrari-f430-2930661_960_720.jpg",
			},
		},
		{
			ID:         uuid.MustParse("6a5011a1-fe1f-47df-9a32-b5346b289391"),
			Status:     models.Live,
			Seller:     "bob",
			AuctionEnd: time.Now().Add(19 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Audi",
				Model:    "R8",
				Color:    "White",
				Mileage:  10050,
				Year:     2021,
				ImageUrl: "https://cdn.pixabay.com/photo/2019/12/26/20/50/audi-r8-4721217_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("40490065-dac7-46b6-acc4-df507e0d6570"),
			Status:       models.Live,
			ReservePrice: 20000,
			Seller:       "tom",
			AuctionEnd:   time.Now().Add(20 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Audi",
				Model:    "TT",
				Color:    "Black",
				Mileage:  25400,
				Year:     2020,
				ImageUrl: "https://cdn.pixabay.com/photo/2016/09/01/15/06/audi-1636320_960_720.jpg",
			},
		},
		{
			ID:           uuid.MustParse("3659ac24-29dd-407a-81f5-ecfe6f924b9b"),
			Status:       models.Live,
			ReservePrice: 20000,
			Seller:       "bob",
			AuctionEnd:   time.Now().Add(48 * 24 * time.Hour),
			Item: models.Item{
				Make:     "Ford",
				Model:    "Model T",
				Color:    "Rust",
				Mileage:  150150,
				Year:     1938,
				ImageUrl: "https://cdn.pixabay.com/photo/2017/08/02/19/47/vintage-2573090_960_720.jpg",
			},
		},
	}

	result := inits.DB.Create(auctions)
	slog.Info("Created " + strconv.FormatInt(result.RowsAffected, 10) + " records.")
}

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

	seedData()
}
