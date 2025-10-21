package models

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateItemIndexes(db *mongo.Database) {
	index := mongo.IndexModel{
		Keys: bson.D{
			{Key: "make", Value: "text"},
			{Key: "model", Value: "text"},
			{Key: "color", Value: "text"},
		},
	}

	_, err := db.Collection("items").Indexes().CreateOne(context.Background(), index)
	if err != nil {
		slog.Error("Unable to create indexes: " + err.Error())
	} else {
		slog.Debug("Item Indexes created")
	}
}
