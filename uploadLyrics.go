package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Joshua1227/highwayLyricsApi/database"
	"github.com/Joshua1227/highwayLyricsApi/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholder with your Atlas connection string

const uri = `mongodb+srv://sandeepjoshuadaniel:%s@lyricsdb0.1rri3.mongodb.net/?retryWrites=true&w=majority&appName=LyricsDB0`

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	db_password := database.GetDbCreds("credentials.json")
	opts := options.Client().ApplyURI(fmt.Sprintf(uri, db_password)).SetServerAPIOptions(serverAPI)
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

	file, err := os.Open("Songs 2024.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	count := 0
	var songs []interface{}
	var curSong models.Song
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 5 && strings.Contains(line[:5], ".") && strings.Count(line, ".") == 1 {
			if curSong.Lyrics != "" {
				songs = append(songs, curSong)
				curSong = initializeSong()
			}
			count++
			curSong.Title = strings.TrimSpace(strings.Split(line, ".")[1])
			curSong.AddedBy = "Highway Lyrics Script"
		} else {
			curSong.Lyrics = fmt.Sprintln(curSong.Lyrics, line)
		}
	}
	songs = append(songs, curSong)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	coll := client.Database("Highway").Collection("lyrics")
	result, err := coll.InsertMany(context.TODO(), songs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))

	// result1, err1 := coll.DeleteMany(context.TODO(), bson.D{})
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }
	// fmt.Println(result1.DeletedCount)

	// for _, id := range result1.InsertedIDs {
	// 	fmt.Printf("Inserted document with _id: %v\n", id)
	// }

	fmt.Println("Total number of songs =", count)
}

func initializeSong() models.Song {
	var song models.Song
	return song
}
