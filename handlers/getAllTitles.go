package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Joshua1227/highwayLyricsApi/database"
	"github.com/Joshua1227/highwayLyricsApi/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTitles(c *gin.Context) {

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
	if err != nil {
		panic(err)
	}

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var data []models.Title

	for cursor.Next(context.TODO()) {
		var result models.Title
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		data = append(data, result)
		fmt.Printf("%+v\n", result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		gin.Logger()
	}

	c.IndentedJSON(http.StatusOK, data)
}
