package models

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitDb(db *mongo.Database) error {
	collection := db.Collection("items")
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	if count == 0 {
		slog.Info("Loading test data.")
		jsonFile, err := os.Open("data/auctions.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var items []Item
		json.Unmarshal(byteValue, &items)
		slog.Info("Loaded " + strconv.Itoa(len(items)) + " items from JSON.")

		result, err := collection.InsertMany(context.TODO(), items)
		if err != nil {
			return err
		}
		slog.Info("Inserted " + strconv.Itoa(len(result.InsertedIDs)) + " items into DB.")
	}
	return nil
}
