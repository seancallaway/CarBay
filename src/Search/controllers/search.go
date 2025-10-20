package controllers

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/seancallaway/CarBay/search/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Search(ctx *gin.Context) {
	db, exists := ctx.Get("DB")
	collection := db.(*mongo.Database).Collection("items")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to connect to database"})
		return
	}

	queryStr := ctx.Query("query")
	var filter bson.D
	if queryStr == "" {
		// No query provided
		filter = bson.D{}
	} else {
		// TODO: Implement filtering
		filter = bson.D{}
	}
	pageNumber, err := strconv.Atoi(ctx.DefaultQuery("pageNumber", "1"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "4"))
	if err != nil || pageSize > 30 {
		pageSize = 4
	}

	totalCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "make", Value: 1}})
	// TODO: See about cursor-based/keyset pagination to improve performance.
	findOptions.SetSkip((int64(pageNumber) - 1) * int64(pageSize))
	findOptions.SetLimit(int64(pageSize))

	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		panic(err)
	}

	var results []models.Item
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results":    results,
		"pageCount":  int(math.Ceil(float64(totalCount) / float64(pageSize))),
		"totalCount": totalCount,
	})
}
