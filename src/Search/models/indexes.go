package models

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateItemIndexes(db *mongo.Database) {
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "make", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "model", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "color", Value: 1}},
		},
	}

	_, err := db.Collection("items").Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		slog.Error("Unable to create indexes: " + err.Error())
	} else {
		slog.Debug("Item Indexes created")
	}
}
