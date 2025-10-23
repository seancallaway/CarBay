package database

import (
	"fmt"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/seancallaway/CarBay/search/internal/utils"
)

func ConnectDB() *mongo.Client {
	host := utils.Getenv("MONGO_HOST", "mongo")
	port := utils.Getenv("MONGO_PORT", "27017")
	user := utils.Getenv("MONGO_USER", "root")
	pass := os.Getenv("MONGO_PASSWORD")
	if pass == "" {
		slog.Error("The 'MONGO_PASSWORD' environment variable must be set.")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, pass, host, port)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}
