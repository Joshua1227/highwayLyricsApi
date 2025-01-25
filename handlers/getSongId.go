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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSongById(c *gin.Context) {

	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	fmt.Println("objectId", objectId)

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

	filter := bson.M{"_id": bson.M{"$eq": objectId}}

	result := coll.FindOne(context.Background(), filter)

	song := models.Song{}
	fmt.Println("result", result)
	decodeErr := result.Decode(&song)

	if decodeErr != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, song)
}
