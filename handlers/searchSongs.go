package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Joshua1227/highwayLyricsApi/database"
	"github.com/Joshua1227/highwayLyricsApi/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchSongs(c *gin.Context) {

	key := c.Param("key")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	db_password := database.GetDbCreds("credentials.json")
	opts := options.Client().ApplyURI(fmt.Sprintf(database.Uri, db_password)).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("Highway").Collection("lyrics")

	key = fmt.Sprintf("\"%s\"", key)
	pipeline := []bson.D{
		{{Key: "$match", Value: bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: key}}}}}},
		{{Key: "$addFields", Value: bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}}},
		{{Key: "$match", Value: bson.D{{Key: "score", Value: bson.D{{Key: "$gte", Value: 0.5}}}}}}, // Minimum score of 0.5
		{{Key: "$sort", Value: bson.D{{Key: "score", Value: -1}}}},                                 // Sort by score (optional)
	}

	cursor, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		panic(err)
	}

	var searchedSongs []models.Song
	// var results []interface{}
	if err = cursor.All(context.TODO(), &searchedSongs); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, searchedSongs)
}
