package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholder with your Atlas connection string

type Creds struct {
	DbPassword string `json:"mongodb"`
}

const uri = `mongodb+srv://sandeepjoshuadaniel:%s@lyricsdb0.1rri3.mongodb.net/?retryWrites=true&w=majority&appName=LyricsDB0`

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	db_password := getDbCreds("credentials.json")
	fmt.Println(fmt.Sprintf(uri, db_password))
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
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	file, err := os.Open("Songbook 2019.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if len(line) > 0 && unicode.IsDigit(rune(line[0])) {
			fmt.Println(line)
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total number of songs = ", count)

}

func getDbCreds(fileName string) string {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var creds Creds

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &creds)

	return creds.DbPassword
}
