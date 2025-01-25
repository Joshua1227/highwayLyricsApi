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

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: key}}}}
	sort := bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "title", Value: 1}, {Key: "lyrics", Value: 1}, {Key: "approvedby", Value: 1}, {Key: "addedby", Value: 1}, {Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}
	seachOpts := options.Find().SetSort(sort).SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, seachOpts)
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
