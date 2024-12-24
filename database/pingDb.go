package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uriTest = `mongodb+srv://sandeepjoshuadaniel:%s@lyricsdb0.1rri3.mongodb.net/?retryWrites=true&w=majority&appName=LyricsDB0`

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	db_password := GetDbCreds("credentials.json")
	opts := options.Client().ApplyURI(fmt.Sprintf(uriTest, db_password)).SetServerAPIOptions(serverAPI)
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
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
